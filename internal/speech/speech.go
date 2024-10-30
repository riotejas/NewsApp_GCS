package speech

import (
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"context"
)

// https://texttospeech.googleapis.com/v1beta1/text:synthesize

type Speech struct {
	Client *texttospeech.Client
	ctx    context.Context
}

func NewSpeechClient() (*Speech, error) {
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Speech{Client: client, ctx: ctx}, nil
}

func (s Speech) SpeechClient(text string) ([]byte, error) {
	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		// Build the voice request
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_FEMALE,
		},
		// Select the type of audio file
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := s.Client.SynthesizeSpeech(s.ctx, &req)
	if err != nil {
		return nil, err
	}

	// The resp's AudioContent is binary.
	//filename := "output.mp3"
	//err = os.WriteFile(filename, resp.AudioContent, 0644)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Printf("Audio content written to file: %v\n", filename)
	return resp.AudioContent, nil
}
