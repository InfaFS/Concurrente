Procedure Ej2 is

----Lectores
Task type Lector;
arrLectores: array(1..L) of Lector;


----Escritores
Task type Escritor;

arrEscritores: array(1..L) of Escritor;

---Admin
Task Admin is
    Entry SolicitudEscritor();
    Entry SolicitudLector();
    Entry SalidaEscritor();
    Entry SalidaLector();

End Admin;


admin: AdministradorDeDB;

Task Body Escritor is

Begin
    loop;
        SELECT 
            Admin.SolicitudEscritor();
            --accede a la base de datos
            Admin.SalidaEscritor();
        else
            DELAY 60.0;
        END SELECT;
    end loop;


end Escritor;

Task Body Lector is

Begin
    loop;
        SELECT 
            Admin.SolicitudLector();
            --accede a la base de datos
            Admin.SalidaLector();
        DELAY 120.0;
            DELAY 300.0;
        END SELECT;
    end loop;
end Lector;


Task Body Admin is 
    CWriters,CReaders: int;
Begin
    CWriters := 0;
    CReaders := 0;
    loop
        SELECT
        when (CReaders == 0) =>
            accept SolicitudEscritor();
            --Espera a que termine de escribir
            accept SalidaEscritor();    
        OR
            when (SolicitudEscritor'count = 0) => --no seria necesario tener una cantidad de escritores porque solo va a haber uno en la BD
            --al mismo tiempo, y si entra aca es porque no hay ningun wirter actualmente ni pidiendo ni escribiendo
                accept SolicitudLector();
                CReaders++;
        OR
            accept SalidaLector();
            CReaders--;
        END SELECT;
    end loop;
end Admin;

Begin


End EJ2;
