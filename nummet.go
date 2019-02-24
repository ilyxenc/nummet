package nummet

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

//структура Вектор
type Vector struct {
	I    int
	J    int
	Data map[int][]float64
}

func CreateVect(size int, method matrMethods) *Vector {
	data := method(size, 1)
	return &Vector{
		I:    size,
		J: 1,
		Data: data,
	}
}

//норма Вектора 1
func (vect *Vector) VectNorm1() float64 {
	var max float64 = vect.Data[0][0]
	for i := 0; i < vect.I; i++ {
		if vect.Data[i][0] > max {
			max = vect.Data[i][0]
		}
	}
	return max
}

//норма Вектора 2
func (vect *Vector) VectNorm2() float64 {
	var sum float64 = 0
	for i := 0; i < vect.I; i++ {
		sum += vect.Data[i][0]
	}
	return sum
}

//норма Вектора 3
func (vect *Vector) VectNorm3(accuracy float64) float64 {
	var sum float64 = 0
	for i := 0; i < vect.I; i++ {
		sum += vect.Data[i][0] * vect.Data[i][0]
	}
	//точность нормы 3
	var accur float64
	accur = math.Pow(10, accuracy)
	return math.Round(math.Sqrt(sum)*accur) / accur
}

//структура Матрица
type Matrix struct {
	I     int //строки
	J     int //столбцы
	Data  map[int][]float64
	DataT map[int][]float64
}

//создание матрицы
func CreateMatr(i, j int, method matrMethods) *Matrix {
	dataT := make(map[int][]float64)
	data := method(i, j)
	for x := 0; x < j; x++ {
		for y := 0; y < i; y++ {
			dataT[x] = append(dataT[x], data[y][x])
		}
	}
	fmt.Println("Матрица", i, "*", j, "задана")
	return &Matrix{
		I:     i,
		J:     j,
		Data:  data,
		DataT: dataT,
	}
}

//вывод Матрицы
func (matr *Matrix) MatrixOutput(typeMatr string) {
	var ii int
	var jj int
	var data map[int][]float64
	if strings.ToLower(typeMatr) == "o" { //Оригинальная матрица
		ii = matr.I
		jj = matr.J
		data = matr.Data
	} else if strings.ToLower(typeMatr) == "t" { //Транспанированная матрица
		ii = matr.J
		jj = matr.I
		data = matr.DataT
	} else {
		fmt.Println("Матрица выбрана неверно")
	}
	arrJ := make([]float64, jj)
	for i := 0; i < jj; i++ {
		arrJ[i] = float64(i) + 1
	}
	fmt.Printf("%6s", " ")
	fmt.Printf("%8.f", arrJ)
	fmt.Println("\n")
	for i := 0; i < ii; i++ {
		fmt.Printf("%5d ", i+1)
		fmt.Printf("%8.3f", data[i])
		fmt.Println()
	}
}

//норма Матрицы 1 (max i), 2  (max j)
func (matr *Matrix) MatrNorm1(norm string) float64 {
	var ii int
	var jj int
	var data map[int][]float64
	if strings.ToLower(norm) == "1" { //max i
		ii = matr.I
		jj = matr.J
		data = matr.Data
	} else if strings.ToLower(norm) == "2" { //max j
		ii = matr.J
		jj = matr.I
		data = matr.DataT
	} else {
		fmt.Println("Error")
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
func MatrMultiply(a, b *Matrix) *Matrix {
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

//создание синонима
type matrMethods func(int, int) map[int][]float64

//ввод Матрицы с клавиатуры
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

//ввод Матрицы с Rand
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

//ввод Вектора с клавиатуры
func KeyVect(size, j int) map[int][]float64 {
	return KeyMatr(size, j)
}

//ввод Вектора с Rand
func RandVect(size, j int) map[int][]float64 {
	return RandMatr(size, j)
}

//решение СЛАУ методом Гаусса с выбором главного элемента
func GaussMain(A, B *Matrix, accuracy float64) *Matrix {
	AB := (*A)
	AB.J = (*A).J + 1
	for i := 0; i < (*A).I; i++ {
		AB.Data[i] = append(AB.Data[i], (*B).Data[i][0])
	}
	fmt.Println(AB)
	var n = AB.I //число уравнений
	//MAX в строке
	for i := 0; i < n; i++ {
		max := math.Abs(AB.Data[i][i])
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(AB.Data[i][i]) > max{
				max = math.Abs(AB.Data[i][i])
				maxRow = k
			}
		}
		//меняем местами строку с наибольшим элементом
		for k := i; k < n + 1; k++ {
			c := (*A).Data[maxRow][k]
			AB.Data[maxRow][k] = AB.Data[i][k]
			AB.Data[i][k] = c
		}
		//обнуляем все строки ниже текущей
		for k := i + 1; k < n; k++ {
			c := -AB.Data[k][i] / AB.Data[i][i]
			for j := i; j < n + 1; j++ {
				if i == j {
					AB.Data[k][j] = 0
				} else {
					AB.Data[k][j] += c * AB.Data[i][j]
				}
			}
		}
	}

	//решаем для верхнего треугольника
	x := make(map[int][]float64)
	for i := n - 1; i > -1; i-- {
		x[i] = append(x[i], AB.Data[i][n] / AB.Data[i][i])
		for k := i - 1; k > -1; k-- {
			AB.Data[k][n] -= AB.Data[k][i] * x[i][0]
		}
	}

	//округляем результаты
	accur := math.Pow(10, accuracy)
	for i := 0; i < n; i++ {
		x[i][0] = math.Round(x[i][0]*accur) / accur
	}
	return &Matrix{
		I: n,
		Data: x,
	}
}
