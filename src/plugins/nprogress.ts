import NProgress from "nprogress";
import "nprogress/nprogress.css";
import { router } from "./../router/index";

NProgress.configure({ showSpinner: true, minimum: 0.3 });

router.beforeEach(() => {
  NProgress.start();
});

router.afterEach(() => {
  NProgress.done();
});
