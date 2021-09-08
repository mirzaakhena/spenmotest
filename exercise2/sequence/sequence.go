package sequence

type Numbers []int

func (num Numbers) Exist(sequence Numbers) bool {

  chunkLen := len(num) - len(sequence) + 1

  for i := 0; i < chunkLen; i++ {

    chunks := num[i : i+len(sequence)]

    found := true
    for j := 0; j < len(sequence); j++ {

      if sequence[j] != chunks[j] {
        found = false
        break
      }

    }
    if found {
      return true
    }
  }

  return false
}
