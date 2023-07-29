package myPackage

// DogPublic is a public struct because its name starts with a capital letter.
// The fields of a struct are usually private, but they can be made public by starting their name with a capital letter.
// A public struct can be used outside of its package, but a private struct can only be used inside its package.
type DogPublic struct {
	Name string
	Race string
	Size string
}
