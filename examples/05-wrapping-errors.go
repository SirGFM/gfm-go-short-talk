package main

import (
	"encoding/asn1"
	"fmt"

	"github.com/pkg/errors"
)

/* For more information, see the brief explanation about ASN.1 at the end. */
var two_ints = []byte{
	0x02, 0x01, 0x00,
	0x02, 0x01, 0x01,
}

var struct_der = []byte{
	0x30, 0x06,
	0x02, 0x01, 0x00,
	0x02, 0x01, 0x01,
}

func decode(data []byte, val interface{}) error {
	/* Real world example: asn1.Unmarshal returns the
	 * remaining data and an error. If an application
	 * were to consider having any remainder as an
	 * error, people could inadvertedly try to wrap
	 * the error in the second conditional. */

	rest, err := asn1.Unmarshal(data, val)
	if err != nil {
		return errors.Wrap(err, "Failed to decode the ASN.1")
	} else if len(rest) > 0 {
		// err is nil here, so this will fail!!
		fmt.Printf("Remainder: %+v\n", rest)
		return errors.Wrap(err, "Couldn't decode the entire message")
	}

	fmt.Printf("Decoded value: %+v\n", val)
	return nil
}

func main() {
	var bad int

	err := decode(two_ints, &bad)
	fmt.Printf("err for bad: %+v\n", err)
	fmt.Printf("\n----------\n\n")

	var good struct {
		A int
		B int
	}

	err = decode(struct_der, &good)
	fmt.Printf("err for good: %+v\n", err)
}

/* ASN.1 works with a TLV structure:
 *           Tag Len Value
 * Where the Value can be another ASN.1
 * structure.
 *
 * For this example, we have:
 * two_ints:
 *     Tag:   0x02 - Int
 *     Len:   0x01 - 1
 *     Value: 0x00 - 0
 *
 *     Tag:   0x02 - Int
 *     Len:   0x01 - 1
 *     Value: 0x01 - 1
 *
 * (two_ints is not a valid ASN.1,
 *  it is actually two ASN.1 concatenated)
 *
 * struct_der:
 *     Tag:   0x30 - Sequence (compound)
 *     Len:   0x06 - 6
 *     Value: Other ASN.1 values (the same as the
 *                                value of two_ints)
 *
 * (strut_der is a valid ASN.1 because
 *  the two ints are inside a sequence
 *  type)
 *
 * Credits: Vitor Nagata
 */
