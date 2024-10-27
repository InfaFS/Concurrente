Procedure BCRA is


task type APIBanco is
    entry ConsultarDolarOficial( compra,venta: OUT String);
end APIBanco;

task TareaProgramada;

arrAPIsL: array (1..20) of APIBanco;

task body TareaProgramada is
    int c1,c2;
    cotizacions: array(1..2,1..20) of String

Begin
    loop
        for i in 1..20 loop
            select
                arrAPIs(i).ConsultarDolarOficial(c1,c2);
                cotizacions(1,i) := c1;
                cotizacions(1,i) := c2;
            or delay 5.0
                cotizacions(1,i) := " ";
                cotizacions(2,i) := " ";
        end loop;
        mostrar(cotizacions);
    end loop;

end TareaProgramada;


task body APIBanco


Begin   
    loop
        accept ConsultarDolarOficial(compra,venta: OUT string) ConsultarDolarOficial
            compra:= cotizaCompra();
            venta:= cotizaVenta();
        end ConsultarDolarOficial
    end loop;

end APIBanco;


Begin
    Null;
End BRCA;