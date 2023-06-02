package library

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
)

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

// Initiate ChatGPT Function
func InitiateChatGPTFunction(flags Flag) {
	/** Start Chat with ChatGPT */
	if *flags.ChatGPT {
		apiKey := os.Getenv("API_KEY_CHATGPT")
		if apiKey == "" {
			apiKey = *flags.API_KEY
		}
		if apiKey == "" {
			panic("❌ API_KEY_CHATGPT is not set")
		}

		log.SetOutput(new(NullWriter))
		ctx := context.Background()
		client := gpt3.NewClient(apiKey)
		fmt.Print("📟 Ask a question or (quit): ")
		rootCmd := &cobra.Command{
			Use:   "chatgpt",
			Short: "Chat with ChatGPT in console.",
			Run: func(cmd *cobra.Command, args []string) {
				scanner := bufio.NewScanner(os.Stdin)
				quit := false
				prompt := ""

				for !quit {
					if !scanner.Scan() {
						break
					}

					question := scanner.Text()
					switch question {
					case "~":
						GetResponsefromChatGPT(client, ctx, prompt)
						fmt.Print("📟 Ask a question or (quit): ")
						prompt = ""
					case "quit":
						quit = true
					default:
						prompt += question + "\n"
					}
				}
			},
		}
		rootCmd.Execute()
	}
}

// Get Response from ChatGPT
func GetResponsefromChatGPT(client gpt3.Client, ctx context.Context, question string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(13)
	}
	fmt.Printf("\n")
}
