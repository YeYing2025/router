import React from 'react';
import { useTranslation } from 'react-i18next';
import { useLocation } from 'react-router-dom';
import LogsTable from '../../components/LogsTable';
import { AppSection } from '../../router-ui';

const Log = () => {
  const { t } = useTranslation();
  const location = useLocation();
  const isAdminWorkspace = location.pathname.startsWith('/admin/');

  return (
    <div className='dashboard-container'>
      <AppSection
        title={isAdminWorkspace ? t('log.title') : undefined}
      >
        <LogsTable />
      </AppSection>
    </div>
  );
};

export default Log;
