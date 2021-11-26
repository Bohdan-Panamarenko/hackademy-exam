import React from 'react';
import { Navigate, Route } from 'react-router-dom';

interface Props {
  Component: React.FC,
  isLoggedIn: boolean,
  restricted: boolean,
  path: string, 
  exact: true | undefined
}

export const PublicRoute: React.FC<Props> = ({Component, isLoggedIn, restricted, path}) => {
  return (
    <Route path={path} children={() => 
      isLoggedIn && restricted ?
        <Navigate to="/" />
      : <Component />
    } />
  );
}