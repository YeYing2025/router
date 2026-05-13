import React from 'react';
import { AppAlert } from '../../router-ui';

const NotFound = () => (
  <div className='router-not-found'>
    <AppAlert
      type='error'
      title='页面不存在'
      description='请检查你的浏览器地址是否正确'
      showIcon
    />
  </div>
);

export default NotFound;
