package inputs

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"

// 	"github.com/tilt-dev/tilt/internal/tiltfile/starkit"
// )

// func TestInputProps(t *testing.T) {
// 	f := starkit.NewFixture(t, NewPlugin())

// 	f.File("Tiltfile", `
// i = input("foo", "bar")
// print(l.key)
// print(l.value)
// `)

// 	_, err := f.ExecFile("Tiltfile")
// 	require.NoError(t, err)
// 	assert.Equal(t, "foo\nbar", f.PrintOutput())
// }

// func TestInputPropsImmutable(t *testing.T) {
// 	f := starkit.NewFixture(t, NewPlugin())

// 	f.File("Tiltfile", `
// l = link("localhost:4000", "web")
// l.url = "XXX"
// `)

// 	_, err := f.ExecFile("Tiltfile")
// 	require.Error(t, err)
// 	assert.Contains(t, err.Error(), "can't assign to .url field of struct")
// }
