--5
with Ada.Text_IO; use Ada.Text_IO;

procedure Ej6 is

    --Workers le avisan que tienen que realizar un calculo
    Task type Worker;

    --Definicion del admin que avisa
    Task Admin is
        Entry PromedioEntrante(PromedioWorker: IN real );
        Entry RealizarCalculo();
    End Admin;

 

    arrWorkers: array(1..10) of Worker;

    
    --Workers
    Task body Worker is
        arrValores: array (1.. 100000) of real := ?;
        sumaTotal: real := 0;
        resultado: real := 0;
    begin
        Admin.RealizarCalculo()
        for I in 1 .. 100000 loop
            sumaTotal := sumaTotal + arrValores(I);
        end loop
        resultado := sumaTotal / 100000;
        Admin.PromedioTotal(resultado);
    end Worker;



    --Admin
    Task Body Admin is
        PromedioTotal: integer := 0;
    begin
        delay ?;
        for I in 1.. 20 loop
            SELECT
                accept RealizarCalculo();
            OR
                accept PromedioEntrante(PromedioEntrante: IN real) do
                    PromedioTotal := PromedioTotal + PromedioEntrante;
                end PromedioEntrante;
            END SELECT;
        end loop;
    End Admin;

begin
    null;
end Ej6;