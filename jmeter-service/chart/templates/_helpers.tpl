{{/*
Expand the name of the chart.
*/}}
{{- define "jmeter-service.name" -}}
{{- include "common.names.name" . -}}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "jmeter-service.fullname" -}}
{{- include "common.names.fullname" . -}}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "jmeter-service.chart" -}}
{{- include "common.names.chart" . -}}
{{- end }}

{{/*
Common labels
*/}}
{{- define "jmeter-service.labels" -}}
helm.sh/chart: {{ include "jmeter-service.chart" . }}
{{ include "jmeter-service.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}


{{/*
Selector labels
*/}}
{{- define "jmeter-service.selectorLabels" -}}
app.kubernetes.io/name: {{ include "jmeter-service.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "jmeter-service.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "jmeter-service.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
