package magnet

import "testing"

func TestParseMagnetUriLowercaseBase32(t *testing.T) {
	uri := "magnet:?xt=urn:btih:vlaxvsoplmxsg5rsr3kdqun3kipb2lu3&dn=Star.Wars.O.Mandaloriano.e.Grogu.2026.1080p.WEB-DL.x264.DUAL.5.1&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce"

	m, err := ParseMagnetUri(uri)
	if err != nil {
		t.Fatalf("ParseMagnetUri() error = %v", err)
	}

	const wantHash = "aac17ac9cf5b2f2376328ed43851bb521e1d2e9b"
	if got := m.InfoHash.HexString(); got != wantHash {
		t.Fatalf("InfoHash = %q, want %q", got, wantHash)
	}

	const wantTitle = "Star.Wars.O.Mandaloriano.e.Grogu.2026.1080p.WEB-DL.x264.DUAL.5.1"
	if m.DisplayName != wantTitle {
		t.Fatalf("DisplayName = %q, want %q", m.DisplayName, wantTitle)
	}

	if len(m.Trackers) != 1 || m.Trackers[0] != "udp://tracker.openbittorrent.com:80/announce" {
		t.Fatalf("Trackers = %#v", m.Trackers)
	}
}
