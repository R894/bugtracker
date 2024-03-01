export interface Comment {
  id: string
  bugId: string
  content: string
  createdAt: Date
  updatedAt: Date
}

export interface newCommentRequest {
  bugId: string
  content: string
}

export interface updateCommentRequest {
  commentId: string
  content: string
}

export interface deleteCommentRequest {
  commentId: string
}
