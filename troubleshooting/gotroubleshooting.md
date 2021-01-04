# Go Troubleshooting

## Format Specifiers
* In format-strings, **%d** is used as a format specifier for integers (**%x** or **%X** can be used for a hexadecimal representation).
* The **%g** is used for float types (**%f** gives a floating-point, and **%e** gives a scientific notation).
* The **%0nd** shows an integer with n digits, and a leading 0 is necessary.
* The **%n.mg** represents the number with m digits after the decimal sign, and n before it. Instead of g, e and f can also be used. For example, the %5.2e formatting of the value 3.4 gives 3.40e+00.

## Debugging strategy
If you don’t want to use a debugger, the following is useful as a simple debugging strategy:

* Use print-statements (with the fmt.Print functions) at well-chosen places.
* In fmt.Printf functions, use the following specifiers to obtain complete info about variables:
    * **%v** gives us a complete output of the value with its fields.
    * **%+v** gives us a complete output of the value with its fields and the field names.
    * **%#v** gives us a complete output of the value with its fields and qualified type name.
    * **%T** gives us the complete type specification.