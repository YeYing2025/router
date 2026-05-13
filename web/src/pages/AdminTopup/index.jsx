import React from 'react';
import TopupPlansManager from '../../components/TopupPlansManager';
import { AppSection } from '../../router-ui';

const AdminTopup = () => {
  return (
    <div className='dashboard-container'>
      <AppSection>
          <TopupPlansManager />
      </AppSection>
    </div>
  );
};

export default AdminTopup;
