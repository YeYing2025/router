import React from 'react';
import { Spin } from 'antd';

function AppSpin({ description, tip, ...props }) {
  return <Spin {...props} description={description ?? tip} />;
}

export default AppSpin;
