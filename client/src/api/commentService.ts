import { deleteCommentRequest, newCommentRequest, updateCommentRequest } from './models/comment'
import { del, get, post, put } from './utils'

export const createComment = async (commentData: newCommentRequest, token: string) => {
  return post('comments', commentData, token)
}

export const getCommentsByBugId = async(bugId: string, token: string) => {
  return get(`comments/${bugId}`, token)
}

export const updateComment = async (commentData: updateCommentRequest, token: string) => {
  return put('comments', commentData, token)
}

export const deleteComment = async (commentData: deleteCommentRequest, token: string) => {
  return del('comments', commentData, token)
}
