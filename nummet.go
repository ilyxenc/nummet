package nummet

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

//структура Матрица
type Matrix struct {
	I     int
	J     int
	Data  map[int][]float64
	DataT map[int][]float64
}

//создание синонима
type matrMethods func(int, int) map[int][]float64

func CreateMatr(i, j int, method matrMethods) *Matrix {
	data := method(i, j)
	dataT := make(map[int][]float64)
	for x := 0; x < j; x++ {
		for y := 0; y < i; y++ {
			dataT[x] = append(dataT[x], data[y][x])
		}
	}
	return &Matrix{
		I:     i,
		J:     j,
		Data:  data,
		DataT: dataT,
	}
}

func KeyMatr(i, j int) map[int][]float64 {
	data := make(map[int][]float64)
	for y := 0; y < i; y++ {
		for x := 0; x < j; x++ {
			var value float64
			fmt.Scan(&value)
			data[y] = append(data[y], value)
		}
	}
	return data
}

func RandMatr(i, j int) map[int][]float64 {
	data := make(map[int][]float64)
	rand.Seed(time.Now().UTC().UnixNano())
	for y := 0; y < i; y++ {
		for x := 0; x < j; x++ {
			data[y] = append(data[y], float64(rand.Intn(100)-50))
		}
	}
	return data
}

func (matr *Matrix) ShowMatr(typeMatr string) {
	var ii int
	var jj int
	var data map[int][]float64
	if strings.ToLower(typeMatr) == "o" {
		ii = matr.I
		jj = matr.J
		data = matr.Data
	} else if strings.ToLower(typeMatr) == "t" {
		ii = matr.J
		jj = matr.I
		data = matr.DataT
	} else {
		fmt.Println("Матрица выбрана неверно")
	}
	fmt.Println("\nМатрица", ii, "*", jj, ":")
	arrJ := make([]float64, jj)
	for i := 0; i < jj; i++ {
		arrJ[i] = float64(i) + 1
	}
	fmt.Printf("%6s", " ")
	fmt.Printf("%9.f", arrJ)
	fmt.Println("\n")
	for i := 0; i < ii; i++ {
		fmt.Printf("%5d ", i+1)
		fmt.Printf("%9.3f", data[i])
		fmt.Println()
	}
}

//норма Вектора 1
func (matr *Matrix) VectNorm1() float64 {
	if matr.J == 1 {
		var max float64 = matr.Data[0][0]
		for i := 0; i < matr.I; i++ {
			if matr.Data[i][0] > max {
				max = matr.Data[i][0]
			}
		}
		return max
	} else {
		panic("Неверная размерность вектора, число столбцов должно быть 1")
	}
}

//норма Вектора 2
func (matr *Matrix) VectNorm2() float64 {
	if matr.J == 1 {
		var sum float64 = 0
		for i := 0; i < matr.I; i++ {
			sum += matr.Data[i][0]
		}
		return sum
	} else {
		panic("Неверная размерность вектора, число столбцов должно быть 1")
	}
}

//норма Вектора 3
func (matr *Matrix) VectNorm3(accuracy float64) float64 {
	if matr.J == 1 {
		var sum float64 = 0
		for i := 0; i < matr.I; i++ {
			sum += matr.Data[i][0] * matr.Data[i][0]
		}
		//точность нормы 3
		var e float64
		e = math.Pow(10, accuracy)
		return math.Round(math.Sqrt(sum)*e) / e
	} else {
		panic("Неверная размерность вектора, число столбцов должно быть 1")
	}
}

//норма Матрицы 1 (max i), 2  (max j)
func (matr *Matrix) MatrNorm1(norm int) float64 {
	var ii int
	var jj int
	var data map[int][]float64
	if norm == 1 { //max i
		ii = matr.I
		jj = matr.J
		data = matr.Data
	} else if norm == 2 { //max j
		ii = matr.J
		jj = matr.I
		data = matr.DataT
	} else {
		panic("Выбрана несуществующая норма матрицы")
	}
	sum := make([]float64, ii)
	for i := 0; i < ii; i++ {
		sum[i] = 0
		for j := 0; j < jj; j++ {
			sum[i] += data[i][j]
		}
	}
	var sMax float64 = sum[0]
	for i := 0; i < ii; i++ {
		if sum[i] > sMax {
			sMax = sum[i]
		}
	}
	return sMax
}

//перемножение матриц
func MatrMult(a, b *Matrix) *Matrix {
	if (*a).J == (*b).I {
		var prod float64
		c := make(map[int][]float64)
		dataT := make(map[int][]float64)
		for i := 0; i < (*a).I; i++ {
			for j := 0; j < (*b).J; j++ {
				prod = 0
				for k := 0; k < (*a).J; k++ {
					prod += (*a).Data[i][k] * (*b).Data[k][j]
				}
				c[i] = append(c[i], prod)
			}
		}
		for x := 0; x < (*b).J; x++ {
			for y := 0; y < (*a).I; y++ {
				dataT[x] = append(dataT[x], c[y][x])
			}
		}
		return &Matrix{
			I:     (*a).I,
			J:     (*b).J,
			Data:  c,
			DataT: dataT,
		}
	} else {
		panic("Такие матрицы перемножать нельзя, число столбцов матрицы 1 не равно числу строк матрицы 2")
	}
}

//перевод 2 матрицы в одну
func MakeOne(a, b *Matrix) *Matrix {
	var O Matrix = Matrix{I: (*a).I, J: (*a).J + 1, Data: make(map[int][]float64)}
	for i := 0; i < O.I; i++ {
		for j := 0; j < O.J-1; j++ {
			O.Data[i] = append(O.Data[i], (*a).Data[i][j])
		}
		O.Data[i] = append(O.Data[i], (*b).Data[i][0])
	}
	return &Matrix{
		I:    O.I,
		J:    O.J,
		Data: O.Data,
	}
}

//решение СЛАУ методом Гаусса
func Gauss(a, b *Matrix) *Matrix {
	if (*a).J == (*a).I && (*a).J == (*b).I && (*b).J == 1 {
		O := MakeOne(a, b)
		x := make(map[int][]float64)
		n := O.I
		var m float64
		for k := 1; k < n; k++ {
			for j := k; j < n; j++ {
				m = O.Data[j][k-1] / O.Data[k-1][k-1]
				for i := 0; i <= n; i++ {
					O.Data[j][i] = O.Data[j][i] - m*O.Data[k-1][i]
				}
			}
		}
		for i := n - 1; i >= 0; i-- {
			x[i] = append(x[i], O.Data[i][n]/O.Data[i][i])
			for c := n - 1; c > i; c-- {
				x[i][0] -= O.Data[i][c] * x[c][0] / O.Data[i][i]
			}
		}

		return &Matrix{
			I:    O.I,
			J:    1,
			Data: x,
		}
	} else {
		panic("Невозможно решить методом Гаусса")
	}
}

//решение СЛАУ методом Гаусса с выбором главного элемента
func GaussMain(A, B *Matrix, accuracy float64) *Matrix {
	O := MakeOne(A, B)
	x := make(map[int][]float64)
	n := O.I
	//MAX в строке
	for i := 0; i < n; i++ {
		max := math.Abs(O.Data[i][i])
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(O.Data[i][i]) > max {
				max = math.Abs(O.Data[i][i])
				maxRow = k
			}
		}
		//меняем местами строку с наибольшим элементом
		for k := i; k < n+1; k++ {
			c := O.Data[maxRow][k]
			O.Data[maxRow][k] = O.Data[i][k]
			O.Data[i][k] = c
		}
		//обнуляем все строки ниже текущей
		for k := i + 1; k < n; k++ {
			c := -O.Data[k][i] / O.Data[i][i]
			for j := i; j < n+1; j++ {
				if i == j {
					O.Data[k][j] = 0
				} else {
					O.Data[k][j] += c * O.Data[i][j]
				}
			}
		}
	}

	//решаем для верхнего треугольника
	for i := n - 1; i > -1; i-- {
		x[i] = append(x[i], O.Data[i][n]/O.Data[i][i])
		for k := i - 1; k > -1; k-- {
			O.Data[k][n] -= O.Data[k][i] * x[i][0]
		}
	}

	//округляем результаты
	accur := math.Pow(10, accuracy)
	for i := 0; i < n; i++ {
		x[i][0] = math.Round(x[i][0]*accur) / accur
	}
	return &Matrix{
		I:    n,
		J:    1,
		Data: x,
	}
}

//решение СЛАУ методом Жордана-Гаусса
func JordanGauss(a, b *Matrix) *Matrix {
	if (*a).J == (*a).I && (*a).J == (*b).I && (*b).J == 1 {
		O := MakeOne(a, b)
		x := make(map[int][]float64)
		n := O.I

		//прямой ход
		var el float64
		for i := 0; i < n; i++ {
			fl := 0
			for p := i; p < n; p++ {
				for q := i; q < n+1; q++ {
					if O.Data[q][p] != 0 {
						el = O.Data[q][p]
						if q != i {
							for k := 0; k < n+1; k++ {
								c := O.Data[q][k]
								O.Data[q][k] = O.Data[i][k]
								O.Data[i][k] = c
							}
						}
						fl = 1
						break
					}
				}
				if fl == 1 {
					break
				}
			}

			//деление верхней строки на выбранный el
			for p := i; p < n+1; p++ {
				O.Data[i][p] /= el
			}

			for q := 1 + i; q < n; q++ {
				el2 := O.Data[q][i]
				for p := i; p < n+1; p++ {
					O.Data[q][p] -= O.Data[i][p] * el2
				}
			}
		}

		//обратный ход
		fl1 := 0
		for i := n - 1; i > 0; i-- {
			for q := n - 2 - fl1; q > -1; q-- {
				el2 := O.Data[q][i]
				for p := 0; p < n+1; p++ {
					O.Data[q][p] -= O.Data[i][p] * el2
				}
			}
			fl1++
		}

		for i := 0; i < n; i++ {
			x[i] = append(x[i], O.Data[i][n])
		}

		return &Matrix{
			I:    O.I,
			J:    1,
			Data: x,
		}
	} else {
		panic("Невозможно решить методом Гаусса")
	}
}

//метод Якоби
func Jacobi(a, b, c *Matrix, accur float64) *Matrix {
	if (*a).J == (*a).I && (*a).J == (*b).I && (*b).J == 1 {
		O := MakeOne(a, b)
		x := make(map[int][]float64)
		n := O.I
		e := 1 / math.Pow(10, accur)
		fmt.Println(e)
		for i := 0; i < n; i++ {
			x[i] = append(x[i], (*c).Data[i][0])
		}
		var s1 float64 = 0
		var s2 float64 = 0
		for true {
			for i := 0; i < n; i++ {
				if i > 1 {
					for j := 1; j < i-1; j++ {
						s1 += O.Data[i][j] * x[j][0]
					}
				} else {
					s1 = 0
				}
				if i < n {
					for j := i + 1; j < n; j++ {
						s2 += O.Data[i][j] * x[j][0]
					}
				} else {
					s2 = 0
				}
				x[i][0] = (O.Data[i][n] - s1 - s2) / O.Data[i][i]
			}
			if math.Abs(O.Data[0][n]-x[0][0]) < e {
				break
			}
			for i := 0; i < n; i++ {
				O.Data[i][n] = x[i][0]
			}
		}
		for i := 0; i < n; i++ {
			x[i][0] = O.Data[i][n]
		}
		return &Matrix{
			I:    n,
			J:    1,
			Data: x,
		}
	} else {
		panic("Невозможно решить методом Якоби")
	}
}

//метод Зейделя
func Zeidel(a, b, c *Matrix, accur float64) *Matrix {
  if (*a).J == (*a).I && (*a).J == (*b).I && (*b).J == 1 {
		O := MakeOne(a, b)
    x := make(map[int][]float64)
    p := make(map[int][]float64)
		n := O.I
    e := 1 / math.Pow(10,accur)
    for i := 0; i < n; i++ {
      x[i] = append(x[i], (*c).Data[i][0])
      p[i] = append(p[i], float64(0))
		}

    for true {
      for i := 0; i < n; i++ {
        p[i][0] = x[i][0]
      }
      for i := 0; i < n; i++ {
        sum := float64(0)
        for j := 0; j < i; j++ {
          sum += O.Data[i][j] * x[j][0]
        }
        for j := i + 1; j < n; j++ {
          sum += O.Data[i][j] * p[j][0]
        }
        x[i][0] = (O.Data[i][n] - sum) / O.Data[i][i]
      }
      norm := float64(0)
      for i := 0; i < n; i++ {
        norm += (x[i][0] - p[i][0])*(x[i][0] - p[i][0])
      }
      if math.Sqrt(norm) < e {
        break
      }
    }

		return &Matrix{
			I:    O.I,
			J:    1,
			Data: x,
		}
	} else {
		panic("Невозможно решить методом Зейделя")
	}
}
