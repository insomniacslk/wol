package wol

import (
	"fmt"
	"net"
)

func Send(macAddr string) error {
	mac, err := net.ParseMAC(macAddr)
	if err != nil {
		return fmt.Errorf("failed to parse MAC address: %w", err)
	}
	packet := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	for i := 0; i < 16; i++ {
		packet = append(packet, mac...)
	}
	localAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:0")
	if err != nil {
		return fmt.Errorf("failed to resolve local addr: %w", err)
	}
	bcastAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
	if err != nil {
		return fmt.Errorf("failed to resolve broadcast address: %w", err)
	}
	conn, err := net.DialUDP("udp", localAddr, bcastAddr)
	if err != nil {
		return fmt.Errorf("failed to dial UDP addr: %w", err)
	}
	defer conn.Close()
	if _, err := conn.Write(packet); err != nil {
		return fmt.Errorf("failed to write packet: %w", err)
	}
	return nil
}
