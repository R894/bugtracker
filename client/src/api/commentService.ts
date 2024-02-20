import { del, post, put } from './utils'

export const createComment = async (commentData: object) => {
  return post('/comments', commentData)
}

export const updateComment = async (commentData: object) => {
  return put('/comments', commentData)
}

export const deleteComment = async (commentData: object) => {
  return del('/comments', commentData)
}
