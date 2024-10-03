FROM scratch

ADD vsys.empms.web ./

ADD static ./static

ENTRYPOINT [ "/vsys.empms.web" ]