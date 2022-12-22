package main

import (
	"container/list"
	"fmt"
)

// Khai báo kiểu dữ liệu nhân viên
type Employee struct {
	Name   string
	Age    int
	Salary float64
}

// Hàm nhập thông tin nhân viên
func inputEmployee() *Employee {
	var emp Employee
	fmt.Print("Nhập tên nhân viên: ")
	fmt.Scanln(&emp.Name)
	fmt.Print("Nhập tuổi nhân viên: ")
	fmt.Scanln(&emp.Age)
	fmt.Print("Nhập lương nhân viên: ")
	fmt.Scanln(&emp.Salary)
	return &emp
}

// Hàm xuất thông tin nhân viên
func outputEmployee(emp *Employee) {
	fmt.Println("Thông tin nhân viên:")
	fmt.Println("Tên:", emp.Name)
	fmt.Println("Tuổi:", emp.Age)
	fmt.Println("Lương:", emp.Salary)
}
// Hàm tìm kiếm nhân viên trong danh sách liên kết
func findEmployee(employees *list.List, name string) *Employee {
	// Duyệt danh sách và tìm kiếm nhân viên có tên tìm kiếm
	for e := employees.Front(); e != nil; e = e.Next() {
		emp := e.Value.(*Employee)
		if emp.Name == name {
			return emp
		}
	}
	// Trả về nil nếu không tìm thấy nhân viên có tên tìm kiếm
	return nil
}
// Hàm thêm nhân viên vào vị trí sau tên được tìm thấy trong danh sách
func addEmployeeAfter(employees *list.List, name string, emp *Employee) bool {
	// Duyệt danh sách và tìm kiếm nhân viên có tên tìm kiếm
	for e := employees.Front(); e != nil; e = e.Next() {
		if e.Value.(*Employee).Name == name {
			// Thêm nhân viên mới vào vị trí sau nhân viên tìm thấy
			employees.InsertAfter(emp, e)
			return true
		}
	}
	// Trả về false nếu không tìm thấy nhân viên có tên tìm kiếm
	return false
}
// Hàm xóa nhân viên tên được tìm thấy trong danh sách
func deleteEmployee(employees *list.List, name string) bool {
	// Duyệt danh sách và tìm kiếm nhân viên có tên tìm kiếm
	for e := employees.Front(); e != nil; e = e.Next() {
		if e.Value.(*Employee).Name == name {
			// Xóa nhân viên tìm thấy khỏi danh sách
			employees.Remove(e)
			return true
		}
	}
	return  false
}
// Hàm sửa thông tin nhân viên tên được tìm thấy trong danh sách
func updateEmployee(employees *list.List, name string) bool {
	// Duyệt danh sách và tìm kiếm nhân viên có tên tìm kiếm
	for e := employees.Front(); e != nil; e = e.Next() {
		if e.Value.(*Employee).Name == name {
			// Nhập và cập nhật thông tin nhân viên tìm thấy
			emp := inputEmployee()
			e.Value = emp
			return true
		}
	}
	// Trả về false nếu không tìm thấy nhân viên có tên tìm kiếm
	return false
}

func main() {
	// Tạo một danh sách liên kết chứa nhân viên
	employees := list.New()

	// Nhập và thêm nhân viên vào danh sách
	fmt.Println("Nhập thông tin nhân viên:")
	emp := inputEmployee()
	employees.PushBack(emp)

	// Nhập và thêm nhân viên khác vào danh sách
	fmt.Println("Nhập thông tin nhân viên khác:")
	emp = inputEmployee()
	employees.PushBack(emp)

	// Duyệt danh sách và xuất thông tin từng nhân viên
	fmt.Println("Danh sách nhân viên:")
	for e := employees.Front(); e != nil; e = e.Next() {
		outputEmployee(e.Value.(*Employee))
	}
	name := "John"
	emp1 := findEmployee(employees, name)
	if emp1 != nil {
		fmt.Println("Tìm thấy nhân viên có tên", name)
		outputEmployee(emp)
	} else {
		fmt.Println("Không tìm thấy nhân viên có tên", name)
	}
/////
	name2 := "John"
	emp2 := inputEmployee()
	if addEmployeeAfter(employees, name2, emp2) {
		fmt.Println("Đã thêm nhân viên vào vị trí sau tên", name)
	} else {
		fmt.Println("Không tìm thấy nhân viên có ")
	}
	//
	name4 := "John"
	if deleteEmployee(employees, name4) {
		fmt.Println("Đã xóa nhân viên có tên", name)
	} else {
		fmt.Println("Không tìm thấy nhân viên có tên", name)
	}
	////
	name5 := "John"
	if updateEmployee(employees, name5) {
		fmt.Println("Đã cập nhật thông tin nhân viên có tên", name)
	} else {
		fmt.Println("Không tìm thấy nhân viên có tên", name)
	}

}
