package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetOpenAIEmbedding(input string) ([]float32, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//create payload
	payload := map[string]interface{}{
		"model": "text-embedding-3-small",
		//input: string to embedding
		"input": input,
	}
	//convert to json
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/embeddings", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	//send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { //request failed
		//return what inside body
		b, _ := io.ReadAll(resp.Body)
		return nil, errors.New(string(b))
	}

	var result struct {
		Data []struct {
			Embedding []float32 `json:"embedding"`
		} `json:"data"`
		// { EXAMPLE JSON STRUCT
		// 	"data":
		// 		{
		// 			"embedding": [
		// 				-0.000123456789,
		// 				-0.000123456789,
		// 				-0.000123456789,
		// 				-0.000123456789,
		// 				-0.000123456789,
		// 			]
		// 		}
		// }
	}
	//decode json of resp.Body to result struct
	json.NewDecoder(resp.Body).Decode(&result)

	if len(result.Data) == 0 { //no embedding returned
		return nil, errors.New("no embedding returned")
	}
	//data[0] because only have 1 embedding
	return result.Data[0].Embedding, nil
}
