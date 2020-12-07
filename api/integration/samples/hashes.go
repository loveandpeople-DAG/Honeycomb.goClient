package samples

import (
	. "github.com/loveandpeople-DAG/goClient/trinary"
	"strings"
)

var Seed = "HZVEINVKVIKGFRAWRTRXWD9JLIYLCQNCXZRBLDETPIQGKZJRYKZXLTV9JNUVBIAHAGUZVIQWIAWDZ9ACW"

func DefaultHashes() Hashes {
	return Hashes{strings.Repeat("A", 81), strings.Repeat("B", 81)}
}
