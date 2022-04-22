package cmd

import "embed"

//go:embed templates/swal_form.html.gotmpl
var swalForm embed.FS

//go:embed templates/form_group.ts.gotmpl
var formGroup embed.FS
