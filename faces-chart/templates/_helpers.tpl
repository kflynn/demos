# partials.select-image builds a Docker image name, with optional defaulting.
#
# params: .root for the root, .source for the source dict, .default (optional)
# for the default dict
{{- define "partials.select-image" -}}
  {{- if .source.image -}}
    {{- .source.image -}}
  {{- else if .source.imageName -}}
    {{- .source.imageName -}}:{{- include "partials.image-tag" . -}}
  {{- else if .default -}}
    {{- include "partials.select-image"
        (dict "root" .root
              "source" .default) -}}
  {{- else -}}
    {{- include "partials.default-image" . -}}
  {{- end -}}
{{- end -}}

# partials.image-tag finds an appropriate Docker tag, with optional defaulting.
#
# params: .root for the root, .source for the source dict, .default (optional)
# for the default dict
{{- define "partials.image-tag" -}}
  {{- if .source.imageTag -}}
    {{- .source.imageTag -}}
  {{- else if (and .default .default.imageTag) -}}
    {{- .default.imageTag -}}
  {{- else if .root.Values.defaultImageTag -}}
    {{- .root.Values.defaultImageTag -}}
  {{- else -}}
    {{- .root.Chart.AppVersion -}}
  {{- end -}}
{{- end -}}

# partials.default-image finds the default Docker image name.
#
# params: .root for the root
{{- define "partials.default-image" -}}
  {{- if .root.Values.defaultImage -}}
    {{- .root.Values.defaultImage -}}
  {{- else -}}
    {{- .root.Values.defaultRegistry -}}/{{- .root.Values.defaultImageName -}}:{{- .root.Values.defaultImageTag -}}
  {{- end -}}
{{- end -}}

# partials.select-key finds a value in a dict, with optional defaulting.
#
# params: .root for the root, .source for the source dict, .key for the key,
# .default (optional) for the default dict
{{- define "partials.select-key" -}}
  {{- $primary := index .source .key -}}
  {{- if $primary -}}
    {{- $primary -}}
  {{- else if .default -}}
    {{- include "partials.select-key"
        (dict "root" .root "source" .default "key" .key) -}}
  {{- else -}}
    {{- $capitalizedKey := camelcase .key -}}
    {{- $defaultKey := printf "default%s" $capitalizedKey -}}
    {{- index .root.Values $defaultKey -}}
  {{- end -}}
{{- end -}}

# partials.select-env finds a value in a dict, with optional defaulting, and
# if present, formats it as an environment variable.
#
# params: .root for the root, .source for the source dict, .key for the key,
# .name for the env name, .default (optional) for the default dict
{{- define "partials.select-env" -}}
  {{- $value := include "partials.select-key" . -}}
  {{- if $value }}
        - name: {{ .name }}
          value: {{ $value | quote }}
  {{ end -}}
{{- end -}}
