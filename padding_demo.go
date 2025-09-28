package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("=== Why Our Struct Had Same Size ===\n")

	// Our original struct - all fields are naturally aligned
	type AdminResponse struct {
		LicenceIssueDate  time.Time // 24 bytes, aligned to 8-byte boundary
		LicenceExpiryDate time.Time // 24 bytes, aligned to 8-byte boundary
		CreatedAt         time.Time // 24 bytes, aligned to 8-byte boundary
		CompanyName       string    // 16 bytes, aligned to 8-byte boundary
		FullName          string    // 16 bytes, aligned to 8-byte boundary
		Email             string    // 16 bytes, aligned to 8-byte boundary
		Phone             string    // 16 bytes, aligned to 8-byte boundary
		Status            string    // 16 bytes, aligned to 8-byte boundary
		ID                int       // 8 bytes, aligned to 8-byte boundary
	}
	// Total: 3*24 + 5*16 + 8 = 72 + 80 + 8 = 160 bytes

	fmt.Printf("AdminResponse size: %d bytes\n", unsafe.Sizeof(AdminResponse{}))

	fmt.Println("\n=== Example Where Ordering DOES Matter ===\n")

	// Bad ordering - will have padding
	type BadStruct struct {
		A bool  // 1 byte
		B int64 // 8 bytes - needs 8-byte alignment, so 7 bytes padding after A
		C bool  // 1 byte
		D int64 // 8 bytes - needs 8-byte alignment, so 7 bytes padding after C
		E bool  // 1 byte + 7 bytes padding at end for struct alignment
	}

	// Good ordering - minimal padding
	type GoodStruct struct {
		B int64 // 8 bytes
		D int64 // 8 bytes
		A bool  // 1 byte
		C bool  // 1 byte
		E bool  // 1 byte + 5 bytes padding at end
	}

	fmt.Printf("Bad struct size:  %d bytes\n", unsafe.Sizeof(BadStruct{}))
	fmt.Printf("Good struct size: %d bytes\n", unsafe.Sizeof(GoodStruct{}))

	fmt.Println("\n=== Memory Layout Visualization ===")

	fmt.Println("\nBad layout:")
	fmt.Println("A[1] + padding[7] + B[8] + C[1] + padding[7] + D[8] + E[1] + padding[7] = 40 bytes")

	fmt.Println("\nGood layout:")
	fmt.Println("B[8] + D[8] + A[1] + C[1] + E[1] + padding[5] = 24 bytes")

	fmt.Println("\n=== Another Example ===")

	type MixedBad struct {
		A byte  // 1 byte
		B int32 // 4 bytes (needs 4-byte alignment, 3 bytes padding)
		C byte  // 1 byte
		D int32 // 4 bytes (needs 4-byte alignment, 3 bytes padding)
	}

	type MixedGood struct {
		B int32 // 4 bytes
		D int32 // 4 bytes
		A byte  // 1 byte
		C byte  // 1 byte + 2 bytes padding
	}

	fmt.Printf("\nMixed bad size:  %d bytes\n", unsafe.Sizeof(MixedBad{}))
	fmt.Printf("Mixed good size: %d bytes\n", unsafe.Sizeof(MixedGood{}))

	fmt.Println("\n=== Why Our AdminResponse Was Already Optimal ===")
	fmt.Println("• All fields (time.Time=24, string=16, int=8) are multiples of 8")
	fmt.Println("• Go aligns to 8-byte boundaries on 64-bit systems")
	fmt.Println("• No padding needed regardless of order")
	fmt.Println("• Our reordering follows best practices but saves no space here")
}
