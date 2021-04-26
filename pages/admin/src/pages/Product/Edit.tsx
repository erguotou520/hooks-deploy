import React from 'react'
import { RouteChildrenProps } from 'react-router'

const EditPage: React.FC<RouteChildrenProps> = ({ match }) => {
  return <div>{match?.isExact}</div>
}

export default EditPage
