package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {

}

func ReadFile(reader Reader) {
	if c, ok := reader.(Closer); ok {
		c.Close()
	} else {
		println("Type isn't converted")
	}
}

func main() {
	file := &File{}
	ReadFile(file)
}
