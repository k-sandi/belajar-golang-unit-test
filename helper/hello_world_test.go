package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name string
		request string
	}{
		{
			name: "Eko",
			request: "Eko",
		},
		{
			name: "Kurniawan",
			request: "Kurniawan",
		},
		{
			name: "EkoKurniawanKhannedy",
			request: "Eko Kurniawan Khannedy",
		},
		{
			name: "KurniaSandi",
			request: "Kurnia Sandi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Eko", func (b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})

	b.Run("Kurniawan", func (b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kurniawan")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}
func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Kurniawa",
			request:  "Kurniawa",
			expected: "Hello Kurniawa",
		},
		{
			name:     "Khannedy",
			request:  "Khannedy",
			expected: "Hello Khannedy",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})

	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello Kurniawa'")
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// error
		// panic("Result is not 'Hello Eko'")
		// t.Fail()
		t.Error("Result must be 'Hello Eko'")
	}

	fmt.Println("TestHelloWorldEko Done")
}
func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")
	if result != "Hello Khannedy" {
		// error
		// panic("Result is not 'Hello Khannedy'")
		// t.FailNow()
		t.Fatal("Result must be 'Hello Khannedy'")
	}

	fmt.Println("TestHelloWorldKhannedy Done")
}
