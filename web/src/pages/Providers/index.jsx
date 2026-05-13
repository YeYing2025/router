import React from 'react';
import ProvidersManager from '../../components/ProvidersManager';
import { AppSection } from '../../router-ui';

const Providers = () => {
  return (
    <div className='dashboard-container'>
      <AppSection>
        <ProvidersManager />
      </AppSection>
    </div>
  );
};

export default Providers;
