import React from 'react';
import { Navigate, Route } from 'react-router-dom';

interface Props {
  Component: React.FC,
  isLoggedIn: boolean,
  path: string
}

export const PrivatRoute: React.FC<Props> = ({Component, isLoggedIn, path}) => {
  return (
    <Route path={path} children={() => 
      isLoggedIn ?
          <Component />
      : <Navigate to="/" />
    }  />
  );
}