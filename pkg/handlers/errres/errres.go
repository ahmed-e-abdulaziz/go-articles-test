package errres

import "net/http"

type ErrorResponse struct {
	err    string
	status int
}

// Article errors start

func ArticleIdNotFoundResponse() ErrorResponse {
	return ErrorResponse{err: "Invalid or no id was supplied for GetArticleById", status: http.StatusBadRequest}
}

func ArticleByIdError(id string) ErrorResponse {
	return ErrorResponse{err: "Encountered an error while getting articles by id: " + id, status: http.StatusBadRequest}
}

func ArticleNotFound(id string) ErrorResponse {
	return ErrorResponse{err: "No article was found for id: " + id, status: http.StatusNotFound}
}

func ArticleGetAllError() ErrorResponse {
	return ErrorResponse{err: "An error occured while getting all articles", status: http.StatusInternalServerError}
}

func ArticleBindingError() ErrorResponse {
	return ErrorResponse{err: "An error occured while parsing the request body as an article", status: http.StatusBadRequest}
}

func ArticleCreationError() ErrorResponse {
	return ErrorResponse{err: "An error occured while creating an article", status: http.StatusInternalServerError}
}

// Article errors end

// Comment errors start

func CommentBindingError() ErrorResponse {
	return ErrorResponse{err: "An error occured while parsing the request body as a comment", status: http.StatusBadRequest}
}

func CommentCreationError() ErrorResponse {
	return ErrorResponse{err: "An error occured while creating a comment", status: http.StatusInternalServerError}
}

func CommentInvalidArticleIdProvidedError() ErrorResponse {
	return ErrorResponse{err: "Invalid article id provided for the comment", status: http.StatusBadRequest}
}
func CommentGetAllByArticleIdError(articleId string) ErrorResponse {
	return ErrorResponse{err: "An error occured while fetching comments for the articleId: " + articleId, status: http.StatusBadRequest}
}

// Comment errors end
