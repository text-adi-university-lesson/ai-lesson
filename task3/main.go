package main

import (
	"fmt"
	"math"
)

var (
	times              = []float64{2.64, 4.66, 1.87, 4.05, 1.73, 5.31, 1.67, 5.96, 0.13, 5.64, 1.52, 4.07, 0.22, 4.79, 0.73}
	limitLearningTimes = 13
)

type Dataset struct {
	inputs  []float64
	outputs []float64
}

func GenerateLearningAndTestingData(data []float64, limit int) (learning []float64, testing []float64) {
	// Згенерувати навчальні дані, обмежені заданим лімітом
	return data[:limit], data[limit:]
}

func generationWeights(count int) []float64 {
	// Згенерувати початкові ваги для нейронної мережі
	var values []float64
	for i := 0; i < count; i++ {
		values = append(values, 1)
	}
	return values
}

func GenerationDataset(data []float64, scope int) []Dataset {
	// Згенерувати датасет для навчання
	result := make([]Dataset, 0)
	for i := 0; i < len(data)-scope; i++ {
		inputSlice := make([]float64, scope)
		copy(inputSlice, data[i:i+scope])
		result = append(result, Dataset{
			inputs:  inputSlice,
			outputs: []float64{data[i+scope]},
		})
	}
	return result
}

func sigmoid(s float64) float64 {
	// сигмоїдна активаційна функція
	return 1.0 / (1.0 + math.Exp(-s)) * 10
}
func getS(inputs, weights []float64, scope int) float64 {
	result := 0.0
	for i := 0; i < scope; i++ {
		result += inputs[i] * weights[i]
	}
	return result
}

func calcError(yPred, yTrue float64) float64 {
	return math.Pow(yPred-yTrue, 2)
}

func LeaningNeuralNetwork(dataset []Dataset, weights []float64, learningRate float64, epochs int, scope int) []float64 {
	for epoch := 0; epoch < epochs; epoch++ {
		totalEpochError := 0.0
		sum_delta_w := make([]float64, scope)
		for _, data := range dataset {
			s := getS(data.inputs, weights, scope)

			yPred := sigmoid(s)

			currentError := calcError(yPred, data.outputs[0])
			totalEpochError += currentError

			// 3.12
			currentErrorStep := make([]float64, 0)
			for i := 0; i < scope; i++ {
				currentErrorI := (yPred - data.outputs[0]) * (math.Exp(-s) / math.Pow(1+math.Exp(-s), 2)) * data.inputs[i]
				currentErrorStep = append(currentErrorStep, currentErrorI)
			}
			// 3.13
			currentWStep := make([]float64, 0)
			for i := 0; i < scope; i++ {
				currentWStep = append(currentWStep, currentErrorStep[i]*-learningRate)
			}
			// 3.13 середне дельта ваг
			for i := 0; i < scope; i++ {
				sum_delta_w[i] += currentWStep[i]
			}
		}
		for i := 0; i < scope; i++ {
			sum_delta_w[i] = sum_delta_w[i] / float64(len(dataset))
		}
		// 3.15
		//fmt.Println("Сумарна помилка", totalEpochError)
		for i := 0; i < scope; i++ {
			weights[i] = weights[i] + sum_delta_w[i]
		}

		mse := totalEpochError / float64(len(dataset))
		if mse < 0.001 { // Поріг точності
			fmt.Printf("Досягнуто бажаної точності на епосі %d! (Помилка: %.5f)\n", epoch, mse)
			break
		}

		if epoch%5000 == 0 {
			fmt.Printf("Епоха %d, Помилка: %.5f\n", epoch, mse)
		}
	}
	return weights
}
func ProgozideData(inputs []float64, expectedResult float64, weight []float64) {
	s := getS(inputs, weight, len(inputs))
	pred := sigmoid(s)
	fmt.Printf("Вхід: %.2v -> Прогноз: %.5f, Очікуване: %.5f (Похибка: %.5f)\n",
		inputs, pred, expectedResult, math.Abs(pred-expectedResult))
}

func main() {
	fmt.Println("Checking logic functions:")
	fmt.Printf("AND(0, 1) = %v\n", logicalAND(0, 1))
	fmt.Printf("OR(0, 1) = %v\n", logicalOR(0, 1))
	fmt.Printf("NOT(1) = %v\n", logicalNOT(1))
	fmt.Printf("XOR(1, 0) = %v\n", logicalXOR(1, 1))
	fmt.Println("Finish checking logic functions:")

	trainingData, testing := GenerateLearningAndTestingData(times, limitLearningTimes)
	scope := 3

	dataset := GenerationDataset(trainingData, scope)
	for i := range dataset {
		fmt.Printf("Dataset %d: inputs=%.5v, outputs=%.5v\n", i, dataset[i].inputs, dataset[i].outputs)
	}

	newWeights := LeaningNeuralNetwork(dataset, generationWeights(scope), 0.005, 100000, scope)

	ProgozideData(times[10:13], testing[0], newWeights)
	ProgozideData(times[11:14], testing[1], newWeights)
}
