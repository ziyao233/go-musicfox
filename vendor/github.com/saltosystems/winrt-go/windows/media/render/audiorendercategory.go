// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package render

type AudioRenderCategory int32

const SignatureAudioRenderCategory string = "enum(Windows.Media.Render.AudioRenderCategory;i4)"

const (
	AudioRenderCategoryOther                  AudioRenderCategory = 0
	AudioRenderCategoryForegroundOnlyMedia    AudioRenderCategory = 1
	AudioRenderCategoryBackgroundCapableMedia AudioRenderCategory = 2
	AudioRenderCategoryCommunications         AudioRenderCategory = 3
	AudioRenderCategoryAlerts                 AudioRenderCategory = 4
	AudioRenderCategorySoundEffects           AudioRenderCategory = 5
	AudioRenderCategoryGameEffects            AudioRenderCategory = 6
	AudioRenderCategoryGameMedia              AudioRenderCategory = 7
	AudioRenderCategoryGameChat               AudioRenderCategory = 8
	AudioRenderCategorySpeech                 AudioRenderCategory = 9
	AudioRenderCategoryMovie                  AudioRenderCategory = 10
	AudioRenderCategoryMedia                  AudioRenderCategory = 11
)
