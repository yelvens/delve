package proc

import (
	"fmt"
)

const (
	_IMAGE_FILE_MACHINE_UNKNOWN     = 0x0
	_IMAGE_FILE_MACHINE_AM33        = 0x1d3
	_IMAGE_FILE_MACHINE_AMD64       = 0x8664
	_IMAGE_FILE_MACHINE_ARM         = 0x1c0
	_IMAGE_FILE_MACHINE_ARMNT       = 0x1c4
	_IMAGE_FILE_MACHINE_ARM64       = 0xaa64
	_IMAGE_FILE_MACHINE_LOONGARCH64 = 0x6264
	_IMAGE_FILE_MACHINE_EBC         = 0xebc
	_IMAGE_FILE_MACHINE_I386        = 0x14c
	_IMAGE_FILE_MACHINE_IA64        = 0x200
	_IMAGE_FILE_MACHINE_M32R        = 0x9041
	_IMAGE_FILE_MACHINE_MIPS16      = 0x266
	_IMAGE_FILE_MACHINE_MIPSFPU     = 0x366
	_IMAGE_FILE_MACHINE_MIPSFPU16   = 0x466
	_IMAGE_FILE_MACHINE_POWERPC     = 0x1f0
	_IMAGE_FILE_MACHINE_POWERPCFP   = 0x1f1
	_IMAGE_FILE_MACHINE_R4000       = 0x166
	_IMAGE_FILE_MACHINE_SH3         = 0x1a2
	_IMAGE_FILE_MACHINE_SH3DSP      = 0x1a3
	_IMAGE_FILE_MACHINE_SH4         = 0x1a6
	_IMAGE_FILE_MACHINE_SH5         = 0x1a8
	_IMAGE_FILE_MACHINE_THUMB       = 0x1c2
	_IMAGE_FILE_MACHINE_WCEMIPSV2   = 0x169
)

type _PEMachine uint16

// PEMachineString map pe machine to name, See $GOROOT/src/debug/pe/pe.go for detail
var _PEMachineString = map[_PEMachine]string{
	_IMAGE_FILE_MACHINE_UNKNOWN:     "unknown",
	_IMAGE_FILE_MACHINE_AM33:        "am33",
	_IMAGE_FILE_MACHINE_AMD64:       "amd64",
	_IMAGE_FILE_MACHINE_ARM:         "arm",
	_IMAGE_FILE_MACHINE_ARMNT:       "armnt",
	_IMAGE_FILE_MACHINE_ARM64:       "arm64",
	_IMAGE_FILE_MACHINE_LOONGARCH64: "loong64",
	_IMAGE_FILE_MACHINE_EBC:         "ebc",
	_IMAGE_FILE_MACHINE_I386:        "i386",
	_IMAGE_FILE_MACHINE_IA64:        "ia64",
	_IMAGE_FILE_MACHINE_M32R:        "m32r",
	_IMAGE_FILE_MACHINE_MIPS16:      "mips16",
	_IMAGE_FILE_MACHINE_MIPSFPU:     "mipsfpu",
	_IMAGE_FILE_MACHINE_MIPSFPU16:   "mipsfpu16",
	_IMAGE_FILE_MACHINE_POWERPC:     "powerpc",
	_IMAGE_FILE_MACHINE_POWERPCFP:   "powerpcfp",
	_IMAGE_FILE_MACHINE_R4000:       "r4000",
	_IMAGE_FILE_MACHINE_SH3:         "sh3",
	_IMAGE_FILE_MACHINE_SH3DSP:      "sh3dsp",
	_IMAGE_FILE_MACHINE_SH4:         "sh4",
	_IMAGE_FILE_MACHINE_SH5:         "sh5",
	_IMAGE_FILE_MACHINE_THUMB:       "thumb",
	_IMAGE_FILE_MACHINE_WCEMIPSV2:   "wcemipsv2",
}

func (m _PEMachine) String() string {
	str, ok := _PEMachineString[m]
	if ok {
		return str
	}
	return fmt.Sprintf("unknown image file machine code %d\n", uint16(m))
}
