import React from 'react';
import PackagesManager from '../../components/PackagesManager';
import { AppSection } from '../../router-ui';

const Package = () => {
  return (
    <div className='dashboard-container'>
      <AppSection>
        <PackagesManager />
      </AppSection>
    </div>
  );
};

export default Package;
