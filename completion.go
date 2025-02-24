// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai

import (
	"context"
	"net/http"

	"github.com/openai/openai-go/internal/apijson"
	"github.com/openai/openai-go/internal/requestconfig"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/packages/resp"
	"github.com/openai/openai-go/packages/ssestream"
	"github.com/openai/openai-go/shared/constant"
)

// CompletionService contains methods and other services that help with interacting
// with the openai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompletionService] method instead.
type CompletionService struct {
	Options []option.RequestOption
}

// NewCompletionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCompletionService(opts ...option.RequestOption) (r CompletionService) {
	r = CompletionService{}
	r.Options = opts
	return
}

// Creates a completion for the provided prompt and parameters.
func (r *CompletionService) New(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (res *Completion, err error) {
	opts = append(r.Options[:], opts...)
	path := "completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a completion for the provided prompt and parameters.
func (r *CompletionService) NewStreaming(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[Completion]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[Completion](ssestream.NewDecoder(raw), err)
}

// Represents a completion response from the API. Note: both the streamed and
// non-streamed response objects share the same shape (unlike the chat endpoint).
type Completion struct {
	// A unique identifier for the completion.
	ID string `json:"id,omitzero,required"`
	// The list of completion choices the model generated for the input prompt.
	Choices []CompletionChoice `json:"choices,omitzero,required"`
	// The Unix timestamp (in seconds) of when the completion was created.
	Created int64 `json:"created,omitzero,required"`
	// The model used for completion.
	Model string `json:"model,omitzero,required"`
	// The object type, which is always "text_completion"
	//
	// This field can be elided, and will be automatically set as "text_completion".
	Object constant.TextCompletion `json:"object,required"`
	// This fingerprint represents the backend configuration that the model runs with.
	//
	// Can be used in conjunction with the `seed` request parameter to understand when
	// backend changes have been made that might impact determinism.
	SystemFingerprint string `json:"system_fingerprint,omitzero"`
	// Usage statistics for the completion request.
	Usage CompletionUsage `json:"usage,omitzero"`
	JSON  struct {
		ID                resp.Field
		Choices           resp.Field
		Created           resp.Field
		Model             resp.Field
		Object            resp.Field
		SystemFingerprint resp.Field
		Usage             resp.Field
		raw               string
	} `json:"-"`
}

func (r Completion) RawJSON() string { return r.JSON.raw }
func (r *Completion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionChoice struct {
	// The reason the model stopped generating tokens. This will be `stop` if the model
	// hit a natural stop point or a provided stop sequence, `length` if the maximum
	// number of tokens specified in the request was reached, or `content_filter` if
	// content was omitted due to a flag from our content filters.
	//
	// Any of "stop", "length", "content_filter"
	FinishReason string                   `json:"finish_reason,omitzero,required"`
	Index        int64                    `json:"index,omitzero,required"`
	Logprobs     CompletionChoiceLogprobs `json:"logprobs,omitzero,required,nullable"`
	Text         string                   `json:"text,omitzero,required"`
	JSON         struct {
		FinishReason resp.Field
		Index        resp.Field
		Logprobs     resp.Field
		Text         resp.Field
		raw          string
	} `json:"-"`
}

func (r CompletionChoice) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reason the model stopped generating tokens. This will be `stop` if the model
// hit a natural stop point or a provided stop sequence, `length` if the maximum
// number of tokens specified in the request was reached, or `content_filter` if
// content was omitted due to a flag from our content filters.
type CompletionChoiceFinishReason = string

const (
	CompletionChoiceFinishReasonStop          CompletionChoiceFinishReason = "stop"
	CompletionChoiceFinishReasonLength        CompletionChoiceFinishReason = "length"
	CompletionChoiceFinishReasonContentFilter CompletionChoiceFinishReason = "content_filter"
)

type CompletionChoiceLogprobs struct {
	TextOffset    []int64              `json:"text_offset,omitzero"`
	TokenLogprobs []float64            `json:"token_logprobs,omitzero"`
	Tokens        []string             `json:"tokens,omitzero"`
	TopLogprobs   []map[string]float64 `json:"top_logprobs,omitzero"`
	JSON          struct {
		TextOffset    resp.Field
		TokenLogprobs resp.Field
		Tokens        resp.Field
		TopLogprobs   resp.Field
		raw           string
	} `json:"-"`
}

func (r CompletionChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage statistics for the completion request.
type CompletionUsage struct {
	// Number of tokens in the generated completion.
	CompletionTokens int64 `json:"completion_tokens,omitzero,required"`
	// Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens,omitzero,required"`
	// Total number of tokens used in the request (prompt + completion).
	TotalTokens int64 `json:"total_tokens,omitzero,required"`
	// Breakdown of tokens used in a completion.
	CompletionTokensDetails CompletionUsageCompletionTokensDetails `json:"completion_tokens_details,omitzero"`
	// Breakdown of tokens used in the prompt.
	PromptTokensDetails CompletionUsagePromptTokensDetails `json:"prompt_tokens_details,omitzero"`
	JSON                struct {
		CompletionTokens        resp.Field
		PromptTokens            resp.Field
		TotalTokens             resp.Field
		CompletionTokensDetails resp.Field
		PromptTokensDetails     resp.Field
		raw                     string
	} `json:"-"`
}

func (r CompletionUsage) RawJSON() string { return r.JSON.raw }
func (r *CompletionUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of tokens used in a completion.
type CompletionUsageCompletionTokensDetails struct {
	// When using Predicted Outputs, the number of tokens in the prediction that
	// appeared in the completion.
	AcceptedPredictionTokens int64 `json:"accepted_prediction_tokens,omitzero"`
	// Audio input tokens generated by the model.
	AudioTokens int64 `json:"audio_tokens,omitzero"`
	// Tokens generated by the model for reasoning.
	ReasoningTokens int64 `json:"reasoning_tokens,omitzero"`
	// When using Predicted Outputs, the number of tokens in the prediction that did
	// not appear in the completion. However, like reasoning tokens, these tokens are
	// still counted in the total completion tokens for purposes of billing, output,
	// and context window limits.
	RejectedPredictionTokens int64 `json:"rejected_prediction_tokens,omitzero"`
	JSON                     struct {
		AcceptedPredictionTokens resp.Field
		AudioTokens              resp.Field
		ReasoningTokens          resp.Field
		RejectedPredictionTokens resp.Field
		raw                      string
	} `json:"-"`
}

func (r CompletionUsageCompletionTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *CompletionUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of tokens used in the prompt.
type CompletionUsagePromptTokensDetails struct {
	// Audio input tokens present in the prompt.
	AudioTokens int64 `json:"audio_tokens,omitzero"`
	// Cached tokens present in the prompt.
	CachedTokens int64 `json:"cached_tokens,omitzero"`
	JSON         struct {
		AudioTokens  resp.Field
		CachedTokens resp.Field
		raw          string
	} `json:"-"`
}

func (r CompletionUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *CompletionUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionNewParams struct {
	// ID of the model to use. You can use the
	// [List models](https://platform.openai.com/docs/api-reference/models/list) API to
	// see all of your available models, or see our
	// [Model overview](https://platform.openai.com/docs/models) for descriptions of
	// them.
	Model string `json:"model,omitzero,required"`
	// The prompt(s) to generate completions for, encoded as a string, array of
	// strings, array of tokens, or array of token arrays.
	//
	// Note that <|endoftext|> is the document separator that the model sees during
	// training, so if a prompt is not specified the model will generate as if from the
	// beginning of a new document.
	Prompt CompletionNewParamsPromptUnion `json:"prompt,omitzero,required"`
	// Generates `best_of` completions server-side and returns the "best" (the one with
	// the highest log probability per token). Results cannot be streamed.
	//
	// When used with `n`, `best_of` controls the number of candidate completions and
	// `n` specifies how many to return – `best_of` must be greater than `n`.
	//
	// **Note:** Because this parameter generates many completions, it can quickly
	// consume your token quota. Use carefully and ensure that you have reasonable
	// settings for `max_tokens` and `stop`.
	BestOf param.Int `json:"best_of,omitzero"`
	// Echo back the prompt in addition to the completion
	Echo param.Bool `json:"echo,omitzero"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their
	// existing frequency in the text so far, decreasing the model's likelihood to
	// repeat the same line verbatim.
	//
	// [See more information about frequency and presence penalties.](https://platform.openai.com/docs/guides/text-generation)
	FrequencyPenalty param.Float `json:"frequency_penalty,omitzero"`
	// Modify the likelihood of specified tokens appearing in the completion.
	//
	// Accepts a JSON object that maps tokens (specified by their token ID in the GPT
	// tokenizer) to an associated bias value from -100 to 100. You can use this
	// [tokenizer tool](/tokenizer?view=bpe) to convert text to token IDs.
	// Mathematically, the bias is added to the logits generated by the model prior to
	// sampling. The exact effect will vary per model, but values between -1 and 1
	// should decrease or increase likelihood of selection; values like -100 or 100
	// should result in a ban or exclusive selection of the relevant token.
	//
	// As an example, you can pass `{"50256": -100}` to prevent the <|endoftext|> token
	// from being generated.
	LogitBias map[string]int64 `json:"logit_bias,omitzero"`
	// Include the log probabilities on the `logprobs` most likely output tokens, as
	// well the chosen tokens. For example, if `logprobs` is 5, the API will return a
	// list of the 5 most likely tokens. The API will always return the `logprob` of
	// the sampled token, so there may be up to `logprobs+1` elements in the response.
	//
	// The maximum value for `logprobs` is 5.
	Logprobs param.Int `json:"logprobs,omitzero"`
	// The maximum number of [tokens](/tokenizer) that can be generated in the
	// completion.
	//
	// The token count of your prompt plus `max_tokens` cannot exceed the model's
	// context length.
	// [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken)
	// for counting tokens.
	MaxTokens param.Int `json:"max_tokens,omitzero"`
	// How many completions to generate for each prompt.
	//
	// **Note:** Because this parameter generates many completions, it can quickly
	// consume your token quota. Use carefully and ensure that you have reasonable
	// settings for `max_tokens` and `stop`.
	N param.Int `json:"n,omitzero"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on
	// whether they appear in the text so far, increasing the model's likelihood to
	// talk about new topics.
	//
	// [See more information about frequency and presence penalties.](https://platform.openai.com/docs/guides/text-generation)
	PresencePenalty param.Float `json:"presence_penalty,omitzero"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same `seed` and parameters should return
	// the same result.
	//
	// Determinism is not guaranteed, and you should refer to the `system_fingerprint`
	// response parameter to monitor changes in the backend.
	Seed param.Int `json:"seed,omitzero"`
	// Up to 4 sequences where the API will stop generating further tokens. The
	// returned text will not contain the stop sequence.
	Stop CompletionNewParamsStopUnion `json:"stop,omitzero"`
	// Options for streaming response. Only set this when you set `stream: true`.
	StreamOptions ChatCompletionStreamOptionsParam `json:"stream_options,omitzero"`
	// The suffix that comes after a completion of inserted text.
	//
	// This parameter is only supported for `gpt-3.5-turbo-instruct`.
	Suffix param.String `json:"suffix,omitzero"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will
	// make the output more random, while lower values like 0.2 will make it more
	// focused and deterministic.
	//
	// We generally recommend altering this or `top_p` but not both.
	Temperature param.Float `json:"temperature,omitzero"`
	// An alternative to sampling with temperature, called nucleus sampling, where the
	// model considers the results of the tokens with top_p probability mass. So 0.1
	// means only the tokens comprising the top 10% probability mass are considered.
	//
	// We generally recommend altering this or `temperature` but not both.
	TopP param.Float `json:"top_p,omitzero"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor
	// and detect abuse.
	// [Learn more](https://platform.openai.com/docs/guides/safety-best-practices#end-user-ids).
	User param.String `json:"user,omitzero"`
	apiobject
}

func (f CompletionNewParams) IsMissing() bool { return param.IsOmitted(f) || f.IsNull() }

func (r CompletionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CompletionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// ID of the model to use. You can use the
// [List models](https://platform.openai.com/docs/api-reference/models/list) API to
// see all of your available models, or see our
// [Model overview](https://platform.openai.com/docs/models) for descriptions of
// them.
type CompletionNewParamsModel = string

const (
	CompletionNewParamsModelGPT3_5TurboInstruct CompletionNewParamsModel = "gpt-3.5-turbo-instruct"
	CompletionNewParamsModelDavinci002          CompletionNewParamsModel = "davinci-002"
	CompletionNewParamsModelBabbage002          CompletionNewParamsModel = "babbage-002"
)

// Only one field can be non-zero
type CompletionNewParamsPromptUnion struct {
	OfString             param.String
	OfArrayOfStrings     []string
	OfArrayOfTokens      []int64
	OfArrayOfTokenArrays [][]int64
	apiunion
}

func (u CompletionNewParamsPromptUnion) IsMissing() bool { return param.IsOmitted(u) || u.IsNull() }

func (u CompletionNewParamsPromptUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[CompletionNewParamsPromptUnion](u.OfString, u.OfArrayOfStrings, u.OfArrayOfTokens, u.OfArrayOfTokenArrays)
}

// Only one field can be non-zero
type CompletionNewParamsStopUnion struct {
	OfString                  param.String
	OfCompletionNewsStopArray []string
	apiunion
}

func (u CompletionNewParamsStopUnion) IsMissing() bool { return param.IsOmitted(u) || u.IsNull() }

func (u CompletionNewParamsStopUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[CompletionNewParamsStopUnion](u.OfString, u.OfCompletionNewsStopArray)
}
