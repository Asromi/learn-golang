package helper

import "fmt"

func fBasicCode() {

	userID := "rOmi"
	println("Hi", userID, ",")

	fmt.Println("=====================================")
	fmt.Println(">> belajar func add, di helper ( 2 + 3 ) ")
	//fmt.Println(Add(2, 3))

	fmt.Println("=====================================")
	fmt.Println(">> belajar query, conn to 20.88")

	//println(helper.math(2, 4))
	// // // inputan float
	// var nip float64
	// print("id: ")
	// fmt.Scanf("%f", &nip)
	// fmt.Println(nip)

	// // // Loop if // // ===================================================================

	// i := 1
	// for i <= 2 {

	// 	if i == 0 {
	// 		println("Zero")
	// 	} else if i == 1 {
	// 		println("One")
	// 	} else if i == 2 {
	// 		println("Two")
	// 	} else if i == 3 {
	// 		println("Three")
	// 	} else if i == 4 {
	// 		println("Four")
	// 	} else if i == 5 {
	// 		println("Five")
	// 	} else {
	// 		println(i)
	// 	}

	// 	i = i + 1
	// }

	// // for case // // ===================================================================

	// for i := 1; i <= 3; i++ {
	// 	switch i {
	// 	case 0:
	// 		println("Zero")
	// 	case 1:
	// 		println("One")
	// 	case 2:
	// 		println("Two")
	// 	case 3:
	// 		println("Three")
	// 	case 4:
	// 		println("Four")
	// 	case 5:
	// 		println("Five")
	// 	default:
	// 		println("Unknown Number")
	// 	}
	// }

	// // // array average // // ===================================================================

	// x := [5]float64{10, 20, 20, 20, 10}

	// var total float64 = 0
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }
	// println(total / float64(len(x)))

	// // // array slices modify // // ===================================================================

	// jeneng := [5]string{
	// 	"Tau ga?", //0
	// 	"Asromi",  //1
	// 	"itu,",    //2
	// 	"Ganteng", //3
	// 	"Banget",  //4
	// }
	// fmt.Println("arr:jeneng,asli-->", jeneng)
	// // // index[0] di array tidak selalu sama dengan index di slice, semua dihitung dari 0 baik array ataupun slice
	// slice1 := jeneng[1:4] //slice hanya rumus, bukan tampungan dari array
	// fmt.Println("slc:slice1[1:4]-->", slice1)

	// slice2 := jeneng[:3] //merubah isi array, akan berpengaruh ke semua slice
	// fmt.Println("slc:slice2[ :3]-->", slice2)

	// slice2[1] = "rOmi"   // merubah value slice2 index 1
	// slice3 := jeneng[1:] //merubah isi slice, akan merubah data array, slice ini akan jadi 1 kesatuan dengan array parent
	// fmt.Println("slc:slice3[1: ]-->", slice3)

	// slice4 := append(slice1, "abis..", "bis..") // fungsi appaend ini akan membuat slice baru, tdk berpengaruh sama array parent
	// fmt.Println("slc:slice4[1+2]-->", slice4)
	// fmt.Println("arr:jeneng,edit-->", jeneng)

	// // // map // // ===================================================================

	// mapX := make(map[string]int)
	// mapX["B"] = 36
	// mapX["C"] = 34
	// fmt.Println(mapX)
	// fmt.Println(mapX["B"])

	// unsur := map[string]string{
	// 	"H":  "Hydrogen",
	// 	"He": "Helium",
	// 	"Li": "Lithium",
	// 	"Be": "Beryllium",
	// 	"B":  "Boron",
	// 	"C":  "Carbon",
	// 	"N":  "Nitrogen",
	// 	"O":  "Oxygen",
	// 	"F":  "Fluorine",
	// 	"Ne": "Neon",
	// }
	// fmt.Println(unsur)
	// fmt.Println(unsur["O"])

	// elements := map[string]map[string]string{
	// 	"H": map[string]string{
	// 		"name":  "Hydrogen",
	// 		"state": "gas",
	// 	},
	// 	"He": map[string]string{
	// 		"name":  "Helium",
	// 		"state": "gas",
	// 	},
	// 	"Li": map[string]string{
	// 		"name":  "Lithium",
	// 		"state": "solid",
	// 	},
	// 	"Be": map[string]string{
	// 		"name":  "Beryllium",
	// 		"state": "solid",
	// 	},
	// 	"B": map[string]string{
	// 		"name":  "Boron",
	// 		"state": "solid",
	// 	},
	// 	"C": map[string]string{
	// 		"name":  "Carbon",
	// 		"state": "solid",
	// 	},
	// 	"N": map[string]string{
	// 		"name":  "Nitrogen",
	// 		"state": "gas",
	// 	},
	// 	"O": map[string]string{
	// 		"name":  "Oxygen",
	// 		"state": "gas",
	// 	},
	// 	"F": map[string]string{
	// 		"name":  "Fluorine",
	// 		"state": "gas",
	// 	},
	// 	"Ne": map[string]string{
	// 		"name":  "Neon",
	// 		"state": "gas",
	// 	},
	// }
	// if el, ok := elements["O"]; ok {
	// 	fmt.Println(el["name"], el["state"])
	// }

	// // ===================================================================

}
