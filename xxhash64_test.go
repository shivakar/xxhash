package xxhash_test

import (
	"testing"

	"github.com/shivakar/xxhash"
	"github.com/stretchr/testify/assert"
)

// Test data
var data = [...]struct {
	input   string
	intHash uint64
	strHash string
}{
	{"", 17241709254077376921, "ef46db3751d8e999"},
	{"a", 15154266338359012955, "d24ec4f1a98c6e5b"},
	{"ab", 7347350983217793633, "65f708ca92d04a61"},
	{"abc", 4952883123889572249, "44bc2cf5ad770999"},
	{"abcd", 15997673941747208908, "de0327b0d25d92cc"},
	{"abcde", 568411279426701291, "07e3670c0c8dc7eb"},
	{"abcdef", 18053520794346263629, "fa8afd82c423144d"},
	{"abcdefg", 1756566643212976685, "1860940e2902822d"},
	{"abcdefgh", 4238821247360054455, "3ad351775b4634b7"},
	{"abcdefghi", 2878261200250560019, "27f1a34fdbb95e13"},
	{"abcdefghij", 15431718392004447154, "d6287a1de5498bb2"},
	{"abcdefghijk", 9317425860326684896, "814e257441cf78e0"},
	{"abcdefghijkl", 5407054947222279347, "4b09b7d3a233d4b3"},
	{"abcdefghijklm", 10613537093487760165, "934adbc0ebc51325"},
	{"abcdefghijklmn", 15451052746170919700, "d66d2a9c05576b14"},
	{"abcdefghijklmno", 3319742962362437736, "2e1218a2b1375068"},
	{"abcdefghijklmnop", 8200634048103437629, "71ce8137ca2dd53d"},
	{"abcdefghijklmnopq", 10371777424410211330, "8feff49d8f62f402"},
	{"abcdefghijklmnopqr", 8044826640893885351, "6fa4f734e2143ba7"},
	{"abcdefghijklmnopqrs", 13356460928919950511, "b95bae7304a854af"},
	{"abcdefghijklmnopqrst", 18216100934840999070, "fccc974985dbdc9e"},
	{"abcdefghijklmnopqrstu", 1147030514091613153, "0feb122ce2f6dbe1"},
	{"abcdefghijklmnopqrstuv", 7146366723549400179, "632cfeac07d58c73"},
	{"abcdefghijklmnopqrstuvw", 14934442522014779562, "cf41cc59032e08aa"},
	{"abcdefghijklmnopqrstuvwx", 859226432292362299, "0bec95e34669983b"},
	{"abcdefghijklmnopqrstuvwxy", 12794926771280486616, "b190b61ba94f20d8"},
	{"abcdefghijklmnopqrstuvwxyz", 14979520437024293724, "cfe1f278fa89835c"},
	{"1", 13237225503670494420, "b7b41276360564d4"},
	{"12", 6080128442901703586, "5460f49adbe7aba2"},
	{"123", 4353148100880623749, "3c697d223fa7e885"},
	{"12345", 14335752410685132726, "c6f2d2dd0ad64fb6"},
	{"123456", 3111357917913400098, "2b2dc38aaa53c322"},
	{"1234567", 15250435807369532249, "d3a46e9108289359"},
	{"123456789", 10139926970967174787, "8cb841db40e6ae83"},
	{"Hello, World!!", 2478871588397104268, "2266b8937637bc8c"},
	{"How are you doing?", 9377945006376102084, "8225274bfc3228c4"},
	{"Discard medicine more than two years old.", 3635545919351933298, "32740dc06f97c972"},
	{"He who has a shady past knows that nice guys finish last.", 2343727645845144928, "208697e054dcc560"},
	{"I wouldn't marry him with a ten foot pole.", 13535952729635593252, "bbd95d4a82689c24"},
	{"Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave", 17461515967836486239, "f253c441f6b47a5f"},
	{"The days of the digital watch are numbered.  -Tom Stoppard", 5362590729855672671, "4a6bbfdb48d9a15f"},
	{"Nepal premier won't resign.", 4513708028829842048, "3ea3e9899e3b3a80"},
	{"For every action there is an equal and opposite government program.", 2243015080680494443, "1f20ca54f5ce156b"},
	{"His money is twice tainted: 'taint yours and 'taint mine.", 10834653289192767395, "965c6bc316da3fa3"},
	{"There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977", 8818883886871487051, "7a62f82bb064224b"},
	{"It's a tiny change to the code and not completely disgusting. - Bob Manchek", 9093601071930493647, "7e32f5feb5874acf"},
	{"size:  a.out:  bad magic", 9287916070484557602, "80e54e74e4dfbf22"},
	{"The major problem is with sendmail.  -Mark Horton", 1914010469571630337, "1a8fee6197321501"},
	{"Give me a rock, paper and scissors and I will move the world.  CCFestoon", 4091155964190771121, "38c6b4ac6e4fc7b1"},
	{"If the enemy is within range, then so are you.", 3204364720526337937, "2c7830bc61c42791"},
	{"It's well we cannot hear the screams/That we create in others' dreams.", 15112545269588140069, "d1ba8bda5dcfa025"},
	{"You remind me of a TV show, but that's all right: I watch it anyway.", 5763003113255514983, "4ffa4ccc3d10b367"},
	{"C is as portable as Stonehedge!!", 14669492906037496783, "cb948213637a4bcf"},
	{"Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley", 16954109249824575268, "eb49188936f54f24"},
	{"The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule", 14472715278762310744, "c8d969ddc5fefc58"},
	{"How can you write a big system without C++?  -Paul Glick", 13996655840875570232, "c23e1c546ec8e438"},
}

func Test_NewXXHash64(t *testing.T) {
	assert := assert.New(t)
	r1 := new(xxhash.XXHash64)
	r1.Reset()
	r2 := xxhash.NewXXHash64()
	r3 := xxhash.NewSeedXXHash64(0)

	assert.Equal(r1, r2)
	assert.Equal(r1, r3)

	// Size and Blocksize
	assert.Equal(r1.Size(), 8)
	assert.Equal(r3.Size(), r1.Size())
	assert.Equal(r1.BlockSize(), 32)
	assert.Equal(r3.BlockSize(), r1.BlockSize())
}

func Test_Hash(t *testing.T) {
	assert := assert.New(t)

	r := xxhash.NewXXHash64()
	for _, s := range data {
		r.Reset()
		r.Write([]byte(s.input))
		assert.Equal(s.intHash, r.Uint64())
		assert.Equal(s.strHash, r.String())
	}
}

func Test_StreamingWrite(t *testing.T) {
	assert := assert.New(t)

	input := "abcdef123456789"

	r1 := xxhash.NewXXHash64()
	r2 := xxhash.NewXXHash64()

	for i, v := range []byte(input) {
		r1.Reset()
		r1.Write([]byte(input[:i+1]))
		r2.Write([]byte(string(v)))

		assert.Equal(r1.String(), r2.String())
	}

}

func Test_WriteEdgeCases(t *testing.T) {
	assert := assert.New(t)

	// Test first write < 32 bytes and second write > 32 bytes
	r1 := xxhash.NewXXHash64()
	r1.Write([]byte("123456"))
	r1.Write([]byte("abcdefghijklmnopqrstuvwxyz"))

	r2 := xxhash.NewXXHash64()
	r2.Write([]byte("123456abcdefghijklmnopqrstuvwxyz"))

	assert.Equal(r1.String(), r2.String())
}
