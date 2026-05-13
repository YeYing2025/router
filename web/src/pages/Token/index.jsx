import React from 'react';
import TokensTable from '../../components/TokensTable';
import { AppSection } from '../../router-ui';

const Token = () => {
  return (
    <div className='dashboard-container'>
      <AppSection>
        <TokensTable />
      </AppSection>
    </div>
  );
};

export default Token;
