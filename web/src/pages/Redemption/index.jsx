import React from 'react';
import RedemptionsTable from '../../components/RedemptionsTable';
import { AppSection } from '../../router-ui';

const Redemption = () => {
  return (
    <div className='dashboard-container'>
      <AppSection>
          <RedemptionsTable />
      </AppSection>
    </div>
  );
};

export default Redemption;
