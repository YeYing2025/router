import React from 'react';
import { AppSpin } from '../router-ui';

const Loading = ({ prompt: name = 'page' }) => {
  return (
    <div className='router-loading-shell'>
      <AppSpin size='large' description={`加载${name}中...`} />
    </div>
  );
};

export default Loading;
