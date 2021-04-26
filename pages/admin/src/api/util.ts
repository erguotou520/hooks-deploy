export function get<T>(url: string) {
  return fetch(url, {
    method: 'GET'
  }).then(resp => {
    if (resp.ok) {
      return resp.json() as Promise<T>
    }
    return Promise.reject()
  })
}

export function post<T>(url: string, data: any, method = 'POST') {
  return fetch(url, {
    method,
    body: data
  }).then(resp => {
    if (resp.ok) {
      return resp.json() as Promise<T>
    }
    return Promise.reject()
  })
}
