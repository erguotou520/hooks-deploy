import React from 'react'
import { GeistProvider, CssBaseline, Page } from '@geist-ui/react'
import RouterComponent from './router'

export default function App() {
  return (
    <GeistProvider>
      <CssBaseline />
      <Page dotBackdrop>
        <Page.Header center>
          <h2>标题</h2>
        </Page.Header>
        <Page.Content>
          <RouterComponent />
        </Page.Content>
        <Page.Footer>页脚</Page.Footer>
      </Page>
    </GeistProvider>
  )
}
