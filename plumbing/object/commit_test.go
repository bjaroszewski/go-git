	"bytes"
	"strings"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"

	i.Close()
}

func (s *SuiteCommit) TestPatch(c *C) {
	from := s.commit(c, plumbing.NewHash("918c48b83bd081e863dbe1b80f8998f058cd8294"))
	to := s.commit(c, plumbing.NewHash("6ecf0ef2c2dffb796033e5a02219af86ec6584e5"))

	patch, err := from.Patch(to)
	c.Assert(err, IsNil)

	buf := bytes.NewBuffer(nil)
	err = patch.Encode(buf)
	c.Assert(err, IsNil)

	c.Assert(buf.String(), Equals, `diff --git a/vendor/foo.go b/vendor/foo.go
new file mode 100644
index 0000000000000000000000000000000000000000..9dea2395f5403188298c1dabe8bdafe562c491e3
--- /dev/null
+++ b/vendor/foo.go
@@ -0,0 +1,7 @@
+package main
+
+import "fmt"
+
+func main() {
+	fmt.Println("Hello, playground")
+}
`)
	c.Assert(buf.String(), Equals, patch.String())

	from = s.commit(c, plumbing.NewHash("b8e471f58bcbca63b07bda20e428190409c2db47"))
	to = s.commit(c, plumbing.NewHash("35e85108805c84807bc66a02d91535e1e24b38b9"))

	patch, err = from.Patch(to)
	c.Assert(err, IsNil)

	buf.Reset()
	err = patch.Encode(buf)
	c.Assert(err, IsNil)

	c.Assert(buf.String(), Equals, `diff --git a/CHANGELOG b/CHANGELOG
deleted file mode 100644
index d3ff53e0564a9f87d8e84b6e28e5060e517008aa..0000000000000000000000000000000000000000
--- a/CHANGELOG
+++ /dev/null
@@ -1 +0,0 @@
-Initial changelog
diff --git a/binary.jpg b/binary.jpg
new file mode 100644
index 0000000000000000000000000000000000000000..d5c0f4ab811897cadf03aec358ae60d21f91c50d
Binary files /dev/null and b/binary.jpg differ
`)

	c.Assert(buf.String(), Equals, patch.String())
			Author:       Signature{Name: "Foo", Email: "foo@example.local", When: ts},
			Committer:    Signature{Name: "Bar", Email: "bar@example.local", When: ts},
			Message:      "Message\n\nFoo\nBar\nWith trailing blank lines\n\n",
			TreeHash:     plumbing.NewHash("f000000000000000000000000000000000000001"),
			ParentHashes: []plumbing.Hash{plumbing.NewHash("f000000000000000000000000000000000000002")},
			TreeHash:  plumbing.NewHash("0000000000000000000000000000000000000003"),
			ParentHashes: []plumbing.Hash{

func (s *SuiteCommit) TestLongCommitMessageSerialization(c *C) {
	encoded := &plumbing.MemoryObject{}
	decoded := &Commit{}
	commit := *s.Commit

	longMessage := "my message: message\n\n" + strings.Repeat("test", 4096) + "\nOK"
	commit.Message = longMessage

	err := commit.Encode(encoded)
	c.Assert(err, IsNil)

	err = decoded.Decode(encoded)
	c.Assert(err, IsNil)
	c.Assert(decoded.Message, Equals, longMessage)
}