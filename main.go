package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

//var name []string
//var rollno []int
//var branch []string
//var uniqueid []int
var studentsCollection *mongo.Collection
var ctx context.Context
var studentResult *mongo.InsertOneResult
var resulte *mongo.UpdateResult
var resultd *mongo.DeleteResult
var err error

func homepage(m echo.Context) error {
	return m.String(http.StatusOK, "Welcome to IITK!\nADDITION OF A STUDENT ENTRY!\n         localhost:9084/add\nDELETION OF A STUDENT ENTRY!\n         localhost:/delete\nEDITING A STUDENT ENTRY!\n         localhost:9084/edit\nFINDING A STUDENT ENTRY!\n         localhost:9084/find")
}
func addition(m echo.Context) error {
	n := m.FormValue("Name:")
	r := m.FormValue("Roll No:")
	b := m.FormValue("Branch:")
	u := m.FormValue("UniqueID:")
	//roll, e := strconv.Atoi(r)
	//unique, er := strconv.Atoi(u)
	//if e != nil || er != nil {
	//	return m.String(http.StatusBadRequest, "Check your roll number and uniqueID inputs.")
	//}
	studentResult, _ = studentsCollection.InsertOne(ctx, bson.D{
		{Key: "Name:", Value: n},
		{Key: "Roll No:", Value: r}, {Key: "Branch:", Value: b}, {Key: "UniqueID:", Value: u},
	})

	//name = append(name, n)
	//rollno = append(rollno, roll)
	//branch = append(branch, b)
	//uniqueid = append(uniqueid, unique)
	return m.String(http.StatusOK, "Successfully added.")
}
func deletion(m echo.Context) error {
	d := m.FormValue("Delete Category:")
	ind := m.FormValue("Index:")
	resultd, err = studentsCollection.DeleteOne(ctx, bson.M{d: ind})
	if err != nil {
		log.Fatal(err)
	}
	return m.String(http.StatusOK, "Deleted")
	//flag := 0
	//if d == "Name" {
	//	for i, k := range name {
	//		if k == ind {
	//			flag = 1
	//			name = append(name[:i], name[i+1:]...)
	//			rollno = append(rollno[:i], rollno[i+1:]...)
	//			branch = append(branch[:i], branch[i+1:]...)
	//			uniqueid = append(uniqueid[:i], uniqueid[i+1:]...)
	//		}
	//	}
	//} else if d == "Roll No" {
	//	roll, e := strconv.Atoi(ind)
	//	if e != nil {
	//		return m.String(http.StatusBadRequest, "You are requested to enter valid roll number.")
	//	}
	//	for i, k := range rollno {
	//		if k == roll {
	//			flag = 1
	//			name = append(name[:i], name[i+1:]...)
	//			rollno = append(rollno[:i], rollno[i+1:]...)
	//			branch = append(branch[:i], branch[i+1:]...)
	//			uniqueid = append(uniqueid[:i], uniqueid[i+1:]...)
	//		}
	//	}
	//} else if d == "Branch" {
	//	for i, k := range branch {
	//		if k == ind {
	//			flag = 1
	//			name = append(name[:i], name[i+1:]...)
	//			rollno = append(rollno[:i], rollno[i+1:]...)
	//			branch = append(branch[:i], branch[i+1:]...)
	//			uniqueid = append(uniqueid[:i], uniqueid[i+1:]...)
	//		}
	//	}
	//} else if d == "UniqueID" {
	//	unique, e := strconv.Atoi(ind)
	//	if e != nil {
	//		return m.String(http.StatusBadRequest, "You are requested to enter valid uniqueID.")
	//	}
	//	for i, k := range uniqueid {
	//		if unique == k {
	//			flag = 1
	//			name = append(name[:i], name[i+1:]...)
	//			rollno = append(rollno[:i], rollno[i+1:]...)
	//			branch = append(branch[:i], branch[i+1:]...)
	//			uniqueid = append(uniqueid[:i], uniqueid[i+1:]...)
	//		}
	//	}
	//}
	//if flag == 0 {
	//	return m.String(http.StatusBadRequest, "Your request terminated!")
	//} else {
	//	return m.String(http.StatusOK, "Successfully deleted.")
	//}
}
func find(m echo.Context) error {
	s := m.FormValue("Search Category:")
	ind := m.FormValue("Index:")
	filterCursor, err := studentsCollection.Find(ctx, bson.M{s: ind})
	if err != nil {
		log.Fatal(err)
	}
	var details []bson.M
	if err = filterCursor.All(ctx, &details); err != nil {
		log.Fatal(err)
	}

	for _, x := range details {
		for k, v := range x {
			m.String(http.StatusOK, fmt.Sprintf("%v%v\n", k, v))
		}
	}
	return m.String(http.StatusOK, "Done!!")
	//return m.String(http.StatusOK,for _, m := range details {
	//	for k, v := range m {
	//		fmt.Println(k, "__", v)
	//	}
	//} )
	//flag := 0
	//if s == "Name" {
	//	for i, k := range name {
	//		if k == ind {
	//			flag = 1
	//			m.String(http.StatusOK, fmt.Sprintf("Name: %v\nRoll No: %v\nBranch: %v\nUniqueID: %v\n", name[i], rollno[i], branch[i], uniqueid[i]))
	//		}
	//	}
	//} else if s == "Roll No" {
	//	roll, e := strconv.Atoi(ind)
	//	if e != nil {
	//		return m.String(http.StatusBadRequest, "You are requested to enter valid roll number.")
	//	}
	//	for i, k := range rollno {
	//		if k == roll {
	//			flag = 1
	//			m.String(http.StatusOK, fmt.Sprintf("Name: %s\nRoll No: %d\nBranch: %s\nUniqueID: %d\n", name[i], rollno[i], branch[i], uniqueid[i]))
	//		}
	//	}
	//} else if s == "Branch" {
	//	for i, k := range branch {
	//		if ind == k {
	//			flag = 1
	//			m.String(http.StatusOK, fmt.Sprintf("Name: %s\nRoll No: %d\nBranch: %s\nUniqueID: %d\n", name[i], rollno[i], branch[i], uniqueid[i]))
	//		}
	//	}
	//} else if s == "UniqueID" {
	//	unique, e := strconv.Atoi(ind)
	//	if e != nil {
	//		return m.String(http.StatusBadRequest, "You are requested to enter valid uniqueID.")
	//	}
	//	for i, k := range uniqueid {
	//		if unique == k {
	//			flag = 1
	//			m.String(http.StatusOK, fmt.Sprintf("Name: %s\nRoll No: %d\nBranch: %s\nUniqueID: %d\n", name[i], rollno[i], branch[i], uniqueid[i]))
	//		}
	//	}
	//}
	//if flag == 0 {
	//	return m.String(http.StatusBadRequest, "Your request terminated!")
	//} else {
	//	return m.String(http.StatusOK, "Thank you for using us!")
	//}
}
func editing(m echo.Context) error {
	ed := m.FormValue("Edit Category:")
	ind := m.FormValue("Index:")
	edit := m.FormValue("New:")
	resulte, err = studentsCollection.UpdateOne(
		ctx,
		bson.M{ed: ind},
		bson.D{
			{"$set", bson.D{{ed, edit}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	//flag := 0
	//if ed == "Name" {
	//	for i, k := range name {
	//		if k == ind {
	//			flag = 1
	//			x := m.FormValue("New Roll No:")
	//			y := m.FormValue("New UniqueID")
	//			x1, err := strconv.Atoi(x)
	//			y1, err2 := strconv.Atoi(y)
	//			if err != nil || err2 != nil {
	//				return m.String(http.StatusBadRequest, "Check your roll number and uniqueID inputs.")
	//			} else {
	//				name[i] = m.FormValue("New Name:")
	//				rollno[i] = x1
	//				branch[i] = m.FormValue("New Branch")
	//				uniqueid[i] = y1
	//			}
	//		}
	//	}
	//} else if ed == "Branch" {
	//	for i, k := range branch {
	//		if k == ind {
	//			flag = 1
	//			x := m.FormValue("New Roll No:")
	//			y := m.FormValue("New UniqueID")
	//			x1, err := strconv.Atoi(x)
	//			y1, err2 := strconv.Atoi(y)
	//			if err != nil || err2 != nil {
	//				return m.String(http.StatusBadRequest, "Check your roll number and uniqueID inputs.")
	//			} else {
	//				name[i] = m.FormValue("New Name:")
	//				rollno[i] = x1
	//				branch[i] = m.FormValue("New Branch")
	//				uniqueid[i] = y1
	//			}
	//		}
	//	}
	//} else if ed == "Roll No" {
	//	roll, e := strconv.Atoi(ind)
	//	if e != nil {
	//		return m.String(http.StatusBadRequest, "You are requested to enter valid roll number.")
	//	}
	//	for i, k := range rollno {
	//		if k == roll {
	//			flag = 1
	//			x := m.FormValue("New Roll No:")
	//			y := m.FormValue("New UniqueID")
	//			x1, err := strconv.Atoi(x)
	//			y1, err2 := strconv.Atoi(y)
	//			if err != nil || err2 != nil {
	//				return m.String(http.StatusBadRequest, "Check your roll number and uniqueID inputs.")
	//			} else {
	//				name[i] = m.FormValue("New Name:")
	//				rollno[i] = x1
	//				branch[i] = m.FormValue("New Branch")
	//				uniqueid[i] = y1
	//			}
	//		}
	//	}
	//} else if ed == "UniqueID" {
	//	unique, e := strconv.Atoi(ind)
	//	if e != nil {
	//		return m.String(http.StatusBadRequest, "You are requested to enter valid uniqueID.")
	//	}
	//	for i, k := range rollno {
	//		if k == unique {
	//			flag = 1
	//			x := m.FormValue("New Roll No:")
	//			y := m.FormValue("New UniqueID")
	//			x1, err := strconv.Atoi(x)
	//			y1, err2 := strconv.Atoi(y)
	//			if err != nil || err2 != nil {
	//				return m.String(http.StatusBadRequest, "Check your roll number and uniqueID inputs.")
	//			} else {
	//				name[i] = m.FormValue("New Name:")
	//				rollno[i] = x1
	//				branch[i] = m.FormValue("New Branch")
	//				uniqueid[i] = y1
	//			}
	//		}
	//	}
	//}
	//if flag == 0 {
	//	return m.String(http.StatusBadRequest, "Your request terminated!")
	//} else {
	//	return m.String(http.StatusOK, "Successfully updated.")
	//}
	return m.String(http.StatusOK, "Updated!")
}
func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://shwetank21:shwetank21@cluster0.859q77r.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 100000*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	SECDatabase := client.Database("SEC")
	studentsCollection = SECDatabase.Collection("students")
	m := echo.New()
	m.GET("/", homepage)
	m.POST("/add", addition)
	m.POST("/delete", deletion)
	m.POST("/find", find)
	m.POST("/edit", editing)
	m.Start(":9084")
}
