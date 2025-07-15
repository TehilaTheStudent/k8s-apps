{{/* Generate a name for the go-server chart */}}
{{- define "go-server.name" -}}
go-server
{{- end -}}

{{/* Generate a fullname for the go-server chart */}}
{{- define "go-server.fullname" -}}
{{ include "go-server.name" . }}
{{- end -}}
