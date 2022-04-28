{{/*Return the proper serivce image name*/}}
{{- define "service.image" -}}
{{/*{{- include "common.images.image" ( dict "imageRoot" .Values.helmservice.image "global" .Values.global.keptn "defaultTag" .Chart.appVersion) -}}*/}}
{{- $registryName := .imageRoot.registry -}}
{{- $repositoryName := .imageRoot.repository -}}
{{- $tag := include "service.tag" (dict "imageRoot" .imageRoot "global" .global "defaultTag" .defaultTag) -}}
{{- if .global }}
    {{- if .global.registry }}
     {{- $registryName = .global.registry -}}
    {{- end -}}
{{- end -}}
{{- if $registryName }}
{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- else -}}
{{- printf "%s:%s" $repositoryName $tag -}}
{{- end -}}
{{- end -}}

{{/*Return the proper serivce image name*/}}
{{- define "service.tag" -}}
{{/*{{- include "common.images.image" ( dict "imageRoot" .Values.helmservice.image "global" .Values.global.keptn "defaultTag" .Chart.appVersion) -}}*/}}
{{- $tag := "" -}}
{{- if .global -}}
{{/*    Set Image Tag: if globally set or at service level set or use default from Chart.yaml*/}}
    {{- if .global.tag -}}
     {{- $tag = .global.tag -}}
    {{- end -}}
    {{- if .imageRoot.tag -}}
     {{- $tag = .imageRoot.tag -}}
    {{- end -}}
    {{- if not $tag }}
     {{- $tag = .defaultTag -}}
    {{- end -}}
{{- end -}}
{{- printf "%s" $tag -}}
{{- end -}}