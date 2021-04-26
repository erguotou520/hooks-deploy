import React from 'react'
import { RouteChildrenProps } from 'react-router'

const CreatePage: React.FC<RouteChildrenProps> = ({ match }) => {
  return <div>{match?.isExact}</div>
}

export default CreatePage
