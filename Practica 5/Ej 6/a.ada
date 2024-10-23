--5
with Ada.Text_IO; use Ada.Text_IO;

procedure Ej6 is

    Task type Worker is
        Entry realizarCalculo();
    End Worker;

    --Definicion del admin que avisa
    Task Admin is
        Entry PromedioEntrante(PromedioWorker: IN real );
    End Admin;

 

    arrWorkers: array(1..20) of Worker;

    --Jugador
    Task Body Jugador is
        IDEquipo: integer := ?;
        SumaMonedas: integer := 0;
    begin 
        arrEquipos(IDEquipo).Llegada(); --Le envian al equipo correspondiente que llegaron y espoeran a que les haya atendido
        arrEquipos(IDEquipo).Empezar(); --Esperan a que les atiendan el empezar
        for I in 1 .. 15 loop
            SumaMonedas := SumaMonedas + Moneda(); --Realiza la suma de las monedas
        end loop
        arrEquipos(IDEquipo).Juntado(SumaMonedas); --Le manda lo juntado al equipo 
        accept EquipoGanador(IDEquipo: IN integer); --Se queda esperando a que le manden el resultado;
    End Jugador;
    
    --Workers
    Task body Worker is
        arrValores: array (1.. 100000) of real := ?;
        sumaTotal: real := 0;
        resultado: real := 0;
    begin
        accept realizarCalculo();
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
        for I in 1 .. 20 loop
            arrWorkers(I).realizarCalculo();
        end loop;

        for I in 1 .. 20 loop
            accept PromedioEntrante (PromedioWorker: IN real) do
                PromedioTotal := PromedioTotal + PromedioWorker;
            end PromedioEntrante;
        end loop;
    End Admin;

begin
    for I in 1..5loop
        arrEquipos(I).RecibirId(I); --LE DA A CADA UNO EL ID
    end loop;
end Ej6;