export interface Project {
  id: number
  name: string
  token: string
  testScript: string
  buildScript: string
  distDir: string
  spa: boolean
  branch: string
  tag: boolean
  tagMatcher: string
  domain: string
  rewrites: {
    from: string
    to: string
  }[]
}
