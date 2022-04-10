// Optional Todo

package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	var drmike1 = "90f2c9c53f66540e67349e0ab83d8cd0"
	var drmike2 = "1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032"
	GenHashMaps("/home/cabox/workspace/course-materials/materials/lab/7/main/Top304Thousand-probable-v2.txt")
	GetSHA(drmike2)
	GetMD5(drmike1)
}
