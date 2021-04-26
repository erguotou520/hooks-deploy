import { Project } from '../types'
import { get, post } from './util'

export async function getProjects() {
  try {
    const data = await get<Project[]>('/api/projects')
    return {
      error: false,
      data
    }
  } catch (error) {
    return {
      error: error
    }
  }
}

export async function getProject(id: number) {
  try {
    const data = await get<Project>(`/api/projects/${id}`)
    return {
      error: false,
      data
    }
  } catch (error) {
    return {
      error: error
    }
  }
}

export async function saveProject(project: Project) {
  try {
    const { id } = await post<{ id: number }>(`/api/projects`, project)
    return {
      error: false,
      data: id
    }
  } catch (error) {
    return {
      error: error
    }
  }
}

export async function deleteProject(id: number) {
  try {
    await post<number>(`/api/projects/${id}`, {}, 'DELETE')
    return {
      error: false,
      data: id
    }
  } catch (error) {
    return {
      error: error
    }
  }
}
