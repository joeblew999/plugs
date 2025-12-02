package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joeblew999/plugs/internal/version"
)

const binaryName = "fakeprinter"

// A minimal fake printer that accepts websocket connections on :8883 with a
// self-signed TLS cert and echoes basic JSON. Useful for testing x1ctl locally.

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	addr := flag.String("addr", ":8883", "listen address for fake printer (TLS websocket)")
	showVersion := flag.Bool("version", false, "print version and exit")
	checkUpdate := flag.Bool("check-update", false, "check for available updates")
	doUpdate := flag.Bool("update", false, "update to latest release from GitHub")
	openDocs := flag.Bool("docs", false, "open documentation in browser")
	openDocsDev := flag.Bool("docs-dev", false, "open developer/technical docs in browser")
	flag.Parse()

	if *showVersion {
		fmt.Println(version.Info())
		os.Exit(0)
	}

	if *openDocs || *openDocsDev {
		docType := version.DocMain
		if *openDocsDev {
			docType = version.DocTech
		}
		if err := version.OpenDocs(binaryName, docType); err != nil {
			log.Fatalf("open docs: %v", err)
		}
		os.Exit(0)
	}

	if *checkUpdate {
		latest, needsUpdate, err := version.CheckUpdate()
		if err != nil {
			log.Fatalf("check update: %v", err)
		}
		if needsUpdate {
			fmt.Printf("Update available: %s (run '%s --update' to upgrade)\n", latest, binaryName)
		} else {
			fmt.Println("You are running the latest version.")
		}
		os.Exit(0)
	}

	if *doUpdate {
		latest, needsUpdate, err := version.CheckUpdate()
		if err != nil {
			log.Fatalf("check update: %v", err)
		}
		if !needsUpdate {
			fmt.Println("Already at latest version:", version.Version)
			os.Exit(0)
		}
		fmt.Printf("Updating from %s to %s...\n", version.Version, latest)
		if err := version.SelfUpdate(binaryName); err != nil {
			log.Fatalf("update failed: %v", err)
		}
		fmt.Println("Update complete. Restart the program to use the new version.")
		os.Exit(0)
	}

	cert, err := selfSignedCert()
	if err != nil {
		log.Fatalf("generate cert: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleWS)

	server := &http.Server{
		Addr:      *addr,
		Handler:   mux,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{*cert}},
	}

	log.Printf("fake printer listening on wss://localhost%s", *addr)
	if err := server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
		log.Fatalf("serve: %v", err)
	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade error: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("client connected: %s", r.RemoteAddr)

	hello := map[string]any{
		"hello": "fake-printer",
		"ts":    time.Now().Unix(),
	}
	if err := conn.WriteJSON(hello); err != nil {
		log.Printf("write hello: %v", err)
		return
	}

	for {
		conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("read: %v", err)
			return
		}
		var msg map[string]any
		_ = json.Unmarshal(data, &msg)
		log.Printf("recv: %s", string(data))

		resp := map[string]any{
			"echo": msg,
			"ts":   time.Now().Unix(),
		}
		if err := conn.WriteJSON(resp); err != nil {
			log.Printf("write: %v", err)
			return
		}
	}
}

func selfSignedCert() (*tls.Certificate, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	serial, err := rand.Int(rand.Reader, big.NewInt(1<<62))
	if err != nil {
		return nil, err
	}

	tmpl := x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			CommonName: "fake-printer.local",
		},
		NotBefore: time.Now().Add(-time.Hour),
		NotAfter:  time.Now().Add(24 * time.Hour),
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		DNSNames: []string{"localhost"},
		IPAddresses: []net.IP{
			net.ParseIP("127.0.0.1"),
			net.ParseIP("::1"),
		},
	}

	der, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &priv.PublicKey, priv)
	if err != nil {
		return nil, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}
