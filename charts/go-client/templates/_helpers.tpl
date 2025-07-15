{{/* Generate a name for the go-client chart */}}
{{- define "go-client.name" -}}
go-client
{{- end -}}

{{/* Generate a fullname for the go-client chart */}}
{{- define "go-client.fullname" -}}
{{ include "go-client.name" . }}
{{- end -}}
